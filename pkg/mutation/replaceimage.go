package mutation

import (
	"github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
)

type replaceimage struct {
	Logger logrus.FieldLogger
}


func (se replaceimage) Name() string {
	return "replaceimage"
}

func (se replaceimage) Mutate(pod *corev1.Pod) (*corev1.Pod, error) {
	se.Logger = se.Logger.WithField("mutation", se.Name())
	mpod := pod.DeepCopy()
	for i, container := range mpod.Spec.Containers {
		se.Logger.Debugf("original image %%s", container.Image)
		if container.Name == "mutateme" {
			mpod.Spec.Containers[i].Image = "ubuntu"
		}
	}

	return mpod, nil
}
