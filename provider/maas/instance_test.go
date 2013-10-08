// Copyright 2013 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package maas

import (
	"fmt"

	gc "launchpad.net/gocheck"

	"launchpad.net/juju-core/instance"
)

type instanceTest struct {
	providerSuite
}

var _ = gc.Suite(&instanceTest{})

func (s *instanceTest) TestId(c *gc.C) {
	jsonValue := `{"system_id": "system_id", "test": "test"}`
	obj := s.testMAASObject.TestServer.NewNode(jsonValue)
	resourceURI, _ := obj.GetField("resource_uri")
	instance := maasInstance{&obj, s.environ}

	c.Check(string(instance.Id()), gc.Equals, resourceURI)
}

func (s *instanceTest) TestString(c *gc.C) {
	jsonValue := `{"hostname": "thethingintheplace", "system_id": "system_id", "test": "test"}`
	obj := s.testMAASObject.TestServer.NewNode(jsonValue)
	instance := &maasInstance{&obj, s.environ}
	hostname, err := instance.DNSName()
	c.Assert(err, gc.IsNil)
	expected := hostname + ":" + string(instance.Id())
	c.Assert(fmt.Sprint(instance), gc.Equals, expected)

	// For good measure, test what happens if we don't have a hostname.
	jsonValue = `{"system_id": "system_id", "test": "test"}`
	obj = s.testMAASObject.TestServer.NewNode(jsonValue)
	instance = &maasInstance{&obj, s.environ}
	hostname, err = instance.DNSName()
	c.Assert(err, gc.NotNil)
	expected = fmt.Sprintf("<DNSName failed: %q>", err) + ":" + string(instance.Id())
	c.Assert(fmt.Sprint(instance), gc.Equals, expected)
}

func (s *instanceTest) TestRefreshInstance(c *gc.C) {
	jsonValue := `{"system_id": "system_id", "test": "test"}`
	obj := s.testMAASObject.TestServer.NewNode(jsonValue)
	s.testMAASObject.TestServer.ChangeNode("system_id", "test2", "test2")
	instance := maasInstance{&obj, s.environ}

	err := instance.refreshInstance()

	c.Check(err, gc.IsNil)
	testField, err := (*instance.maasObject).GetField("test2")
	c.Check(err, gc.IsNil)
	c.Check(testField, gc.Equals, "test2")
}

func (s *instanceTest) TestDNSName(c *gc.C) {
	jsonValue := `{"hostname": "DNS name", "system_id": "system_id"}`
	obj := s.testMAASObject.TestServer.NewNode(jsonValue)
	instance := maasInstance{&obj, s.environ}

	dnsName, err := instance.DNSName()

	c.Check(err, gc.IsNil)
	c.Check(dnsName, gc.Equals, "DNS name")

	// WaitDNSName() currently simply calls DNSName().
	dnsName, err = instance.WaitDNSName()

	c.Check(err, gc.IsNil)
	c.Check(dnsName, gc.Equals, "DNS name")
}

func (s *instanceTest) TestAddresses(c *gc.C) {
	jsonValue := `{
			"hostname": "DNS name", 
			"system_id": "system_id", 
			"ip_addresses": [ "1.2.3.4", "fe80::d806:dbff:fe23:1199" ]
		}`
	obj := s.testMAASObject.TestServer.NewNode(jsonValue)
	inst := maasInstance{&obj, s.environ}

	expected := []instance.Address{
		{
			"DNS name",
			instance.HostName,
			"",
			instance.NetworkPublic,
		},
		{
			"1.2.3.4",
			instance.Ipv4Address,
			"",
			instance.NetworkUnknown,
		},
		{
			"fe80::d806:dbff:fe23:1199",
			instance.Ipv6Address,
			"",
			instance.NetworkUnknown,
		},
	}

	addr, err := inst.Addresses()

	c.Assert(err, gc.IsNil)
	c.Check(addr, gc.DeepEquals, expected)
}

func (s *instanceTest) TestAddressesMissing(c *gc.C) {
	// Older MAAS versions do not have ip_addresses returned, for these
	// just the DNS name should be returned without error.
	jsonValue := `{
		"hostname": "testing.invalid",
		"system_id": "system_id"
		}`
	obj := s.testMAASObject.TestServer.NewNode(jsonValue)
	inst := maasInstance{&obj, s.environ}

	addr, err := inst.Addresses()
	c.Assert(err, gc.IsNil)
	c.Check(addr, gc.DeepEquals, []instance.Address{
		{"testing.invalid", instance.HostName, "", instance.NetworkPublic},
	})
}

func (s *instanceTest) TestAddressesInvalid(c *gc.C) {
	jsonValue := `{
		"hostname": "testing.invalid",
		"system_id": "system_id",
		"ip_addresses": "incompatible"
		}`
	obj := s.testMAASObject.TestServer.NewNode(jsonValue)
	inst := maasInstance{&obj, s.environ}

	_, err := inst.Addresses()
	c.Assert(err, gc.NotNil)
}

func (s *instanceTest) TestAddressesInvalidContents(c *gc.C) {
	jsonValue := `{
		"hostname": "testing.invalid",
		"system_id": "system_id",
		"ip_addresses": [42]
		}`
	obj := s.testMAASObject.TestServer.NewNode(jsonValue)
	inst := maasInstance{&obj, s.environ}

	_, err := inst.Addresses()
	c.Assert(err, gc.NotNil)
}
