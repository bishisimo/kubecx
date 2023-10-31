package kubectx

import (
	"github.com/duke-git/lancet/v2/fileutil"
	"kubecx/pkg/ui"
	"kubecx/util"
)

func (c *KubeCtx) InitMainConfig() error {
	if util.IsFileExists(c.mainContext) {
		if !util.IsRegularFile(c.mainContext) {
			return nil
		}
		if !ui.Confirm("kube config file is init, are you sure cover that?") {
			return nil
		}
	}

	err := fileutil.CopyFile(c.mainContext, c.wardContext)
	if err != nil {
		return err
	}
	if !util.IsFileExists(c.backupContext) {
		fileutil.CopyFile(c.wardContext, c.backupContext)
	}
	return nil
}
