package docker

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"github.com/srepio/cli/internal/driver"
	"github.com/srepio/cli/internal/metadata"
)

type Docker struct {
	client client.APIClient
}

func NewDockerDriver() (*Docker, error) {
	d := &Docker{}

	c, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}
	d.client = c

	return d, nil
}

func (d *Docker) Create(s metadata.Scenario) (driver.Instance, error) {
	instance := Container{
		Name:       s.Name,
		Image:      fmt.Sprintf("%s%s:%s", driver.ImagePrefix, s.Name, s.Version),
		Ports:      s.Ports,
		Volumes:    s.Volumes,
		Privileged: s.Privileged,
	}

	return &instance, nil
}

func (d *Docker) Run(ctx context.Context, i driver.Instance) error {
	c := i.(*Container)

	if err := d.pullImage(ctx, c.Image); err != nil {
		return err
	}
	return d.createContainer(ctx, c)
}

func (d *Docker) ConnectionCommand(i driver.Instance) string {
	c := i.(*Container)

	return fmt.Sprintf("docker exec -it %s bash", c.Name)
}

func (d *Docker) Kill(ctx context.Context, i driver.Instance) error {
	c := i.(*Container)
	err := d.client.ContainerStop(ctx, c.Name, container.StopOptions{})
	if err != nil {
		return err
	}
	return d.client.ContainerRemove(ctx, c.Name, types.ContainerRemoveOptions{})
}

func (d *Docker) Check(ctx context.Context, i driver.Instance) bool {
	c := i.(*Container)
	respID, err := d.client.ContainerExecCreate(ctx, c.Name, types.ExecConfig{
		Cmd:          []string{"/opt/check.sh"},
		AttachStdout: true,
	})
	if err != nil {
		return false
	}

	resp, err := d.client.ContainerExecAttach(ctx, respID.ID, types.ExecStartCheck{})
	if err != nil {
		return false
	}
	defer resp.Close()

	out := new(strings.Builder)
	if _, err := io.Copy(out, resp.Reader); err != nil {
		return false
	}

	// Check it is equal to start, nil *6, text start OK
	return out.String() == string([]byte{1, 0, 0, 0, 0, 0, 0, 2, 79, 75})
}

func (d *Docker) pullImage(ctx context.Context, image string) error {
	reader, err := d.client.ImagePull(ctx, image, types.ImagePullOptions{})
	if err != nil {
		return err
	}
	defer reader.Close()

	buf := new(bytes.Buffer)
	io.Copy(buf, reader)
	return nil
}

func (d *Docker) createContainer(ctx context.Context, c *Container) error {
	resp, err := d.client.ContainerCreate(ctx, d.buildContainerConfig(c), d.buildHostConfig(c), nil, nil, c.Name)
	if err != nil {
		return err
	}
	c.Id = resp.ID

	return d.client.ContainerStart(ctx, c.Id, types.ContainerStartOptions{})
}

func (d *Docker) buildContainerConfig(c *Container) *container.Config {
	ct := &container.Config{
		Image: c.Image,
	}

	ps := nat.PortSet{}
	for _, port := range c.Ports {
		ps[nat.Port(port.Container)] = struct{}{}
	}
	ct.ExposedPorts = ps

	return ct
}

func (d *Docker) buildHostConfig(c *Container) *container.HostConfig {
	hc := &container.HostConfig{
		Privileged: c.Privileged,
	}

	pm := nat.PortMap{}

	for _, port := range c.Ports {
		pm[nat.Port(port.Container)] = []nat.PortBinding{
			{
				HostIP:   "0.0.0.0",
				HostPort: port.Host,
			},
		}
	}
	hc.PortBindings = pm

	vols := []mount.Mount{}

	for _, vol := range c.Volumes {
		vols = append(vols, mount.Mount{
			Type:   mount.TypeBind,
			Source: vol.Host,
			Target: vol.Container,
		})
	}
	hc.Mounts = vols

	return hc
}
