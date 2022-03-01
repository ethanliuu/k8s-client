package deployment

import (
	"context"
	"fmt"
	"github.com/ethanliuuu/k8s-client/common"
	"go.uber.org/zap"
	autoscalingv1 "k8s.io/api/autoscaling/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
	"time"
)

func RestartDeployment(clientset *kubernetes.Clientset, dep string, ns string) (err error) {
	common.LOG.Debug(fmt.Sprintf("RestartDeployment,namespace：%s,deployment：%s",ns,dep))

	patchData := fmt.Sprintf(`{"spec":{"template":{"metadata":{"annotations":{"kubectl.kubernetes.io/restartedAt":"%sT%s+08:00"}}}}}`, time.Now().Format("2006-01-02"), time.Now().Format("15:04:05"))

	_, err = clientset.AppsV1().Deployments(ns).Patch(
		context.TODO(),
		dep,
		types.StrategicMergePatchType,
		[]byte(patchData),
		metav1.PatchOptions{FieldManager: "kubectl-rollout"},
	)
	if err != nil {
		common.LOG.Error("服务重启失败", zap.Any("err: ", err))
		return err
	}
	return nil
}


func ScaleDeployment(clientset *kubernetes.Clientset,ns string, dep string,scaleNum int32) (err error) {
	common.LOG.Debug(fmt.Sprintf("ScaleDeployment,namespace：%s,deployment：%s,scalNumber：%v",ns,dep,scaleNum))
	getScale, err := clientset.AppsV1().Deployments(ns).GetScale(context.TODO(), dep, metav1.GetOptions{})
	if err != nil {
		return err
	}
	scale := autoscalingv1.Scale{
		TypeMeta:   getScale.TypeMeta,
		ObjectMeta: getScale.ObjectMeta,
		Spec:       autoscalingv1.ScaleSpec{Replicas: scaleNum},
		Status:     getScale.Status,
	}
	common.LOG.Debug(fmt.Sprintf("The Deployment has changed form %v to %v",getScale.Spec.Replicas,scaleNum))
	_, err = clientset.AppsV1().Deployments(ns).UpdateScale(context.TODO(), dep, &scale, metav1.UpdateOptions{})
	if err != nil {
		common.LOG.Error("扩缩容失败",zap.Any("err:",err))
		return err
	}
	return nil
}