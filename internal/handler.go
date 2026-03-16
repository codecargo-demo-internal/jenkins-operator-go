package internal

import (
	"context"
	"log"
)

type ReconcileHandler struct {
	Namespace string
}

func (h *ReconcileHandler) Reconcile(ctx context.Context, name string) error {
	log.Printf("Reconciling resource %s in namespace %s", name, h.Namespace)
	return nil
}
