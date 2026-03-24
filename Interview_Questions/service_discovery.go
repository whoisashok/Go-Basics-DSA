package main

import (
	"fmt"
	"math/rand"
)

type Registry struct {
	services map[string][]string
}

func NewRegistry() *Registry {
	return &Registry{
		services: make(map[string][]string),
	}
}

func (r *Registry) Register(name, addr string) {
	r.services[name] = append(r.services[name], addr)
}

func (r *Registry) Discover(name string) []string {
	return r.services[name]
}

type LoadBalancer struct {
	instances []string
	index     int
}

func (lb *LoadBalancer) Next() string {
	addr := lb.instances[lb.index]
	lb.index = (lb.index + 1) % len(lb.instances)
	return addr
}

func ServiceDiscovery() {
	registry := NewRegistry()
	registry.Register("user-service", "http://localhost:8081")
	registry.Register("user-service", "http://localhost:8082")
	registry.Register("order-service", "http://localhost:8082")

	registry.Discover("user-service")  // returns ["http://localhost:8081", "http://localhost:8082"]
	registry.Discover("order-service") // returns ["http://localhost:8082"]

	loadBalancer := LoadBalancer{instances: registry.Discover("user-service")}
	loadBalancer.Next() // returns "http://localhost:8081"

	instances := registry.Discover("user-service")
	fmt.Println("Discovered instances:", instances)

	// simple load balancing (random)
	target := instances[rand.Intn(len(instances))] // randomly returns one of the user-service instances
	fmt.Println("Randomly selected instance:", target)
}
