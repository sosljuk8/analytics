package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type Message struct {
	ent.Schema
}

func (Message) Fields() []ent.Field {
	return []ent.Field{
		field.String("Direction").StructTag(`json:"direction"`),
		field.String("Sent").StructTag(`json:"sent"`),
		field.Int("MailboxId").StructTag(`json:"mailboxId"`),
		field.Int("CrmRid").StructTag(`json:"crmRid"`)}
}
func (Message) Edges() []ent.Edge {
	return nil
}
func (Message) Annotations() []schema.Annotation {
	return nil
}
