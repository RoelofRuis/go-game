package actor

import (
	"github.com/DemonTPx/go-game/lib/actor/property"
	"time"
)

type PhysicsComponent struct {
	BaseComponent
	Velocity   property.Vector3
	friction   float64
	bounciness float64
}

func NewPhysicsComponent(velocity property.Vector3, friction float64, bounciness float64) *PhysicsComponent {
	return &PhysicsComponent{Velocity: velocity, friction: friction, bounciness: bounciness}
}

func (c *PhysicsComponent) Id() ComponentId {
	return Physics
}

func (c *PhysicsComponent) Name() string {
	return "PhysicsComponent"
}

func (c *PhysicsComponent) String() string {
	return "<" + c.Name() + ">"
}

func (c *PhysicsComponent) Update(delta time.Duration) {
	transform := c.owner.GetComponent(Transform)
	if transform == nil {
		return
	}

	t := transform.(*TransformComponent)

	if t.Position.X-t.Scale.X/2 <= 0 && c.Velocity.X < 0 {
		c.Velocity.X = -(c.Velocity.X * c.bounciness)
		t.Position.X = t.Scale.X / 2
	}
	if t.Position.X+t.Scale.X/2 > 1600 && c.Velocity.X > 0 {
		c.Velocity.X = -(c.Velocity.X * c.bounciness)
		t.Position.X = 1600 - t.Scale.X/2
	}

	if t.Position.Y-t.Scale.Y/2 <= 0 && c.Velocity.Y < 0 {
		c.Velocity.Y = -(c.Velocity.Y * c.bounciness)
		t.Position.Y = t.Scale.Y / 2
	}
	if t.Position.Y+t.Scale.Y/2 > 900 && c.Velocity.Y > 0 {
		c.Velocity.Y = -(c.Velocity.Y * c.bounciness)
		t.Position.Y = 900 - t.Scale.Y/2
	}

	c.Velocity.X = c.Velocity.X * (1 - c.friction)
	c.Velocity.Y = c.Velocity.Y * (1 - c.friction)

	t.Position = t.Position.Add(&c.Velocity)
}
