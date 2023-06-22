package osci

import (
	"fmt"
	"strings"

	"github.com/rydyb/telnet"
)

type Client struct {
	telnet.Client
}

func (c *Client) Identity() (string, error) {
	out, err := c.Client.Exec("*idn?")
	if err != nil {
		return "", fmt.Errorf("failed to query identity: %w", err)
	}
	return out, nil
}

func (c *Client) MeasurementList() ([]string, error) {
	out, err := c.Client.Exec("MEASUrement:LIST?")
	if err != nil {
		return nil, fmt.Errorf("failed to query measurement list: %w", err)
	}
	return strings.Split(out, ","), nil
}

func (c *Client) measurementValue(name string) (string, error) {
	out, err := c.Client.Exec(fmt.Sprintf("MEASUrement:%s:VALue?", name))
	if err != nil {
		return "", fmt.Errorf("failed to query measurement value: %w", err)
	}
	return out, nil
}

func (c *Client) Measurements() (map[string]string, error) {
	names, err := c.MeasurementList()
	if err != nil {
		return nil, err
	}

	measurements := make(map[string]string, len(names))
	for _, name := range names {
		value, err := c.measurementValue(name)
		if err != nil {
			return nil, err
		}
		measurements[name] = value
	}
	return measurements, nil
}
