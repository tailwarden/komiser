import { ToastProps } from '@components/toast/Toast';
import useAlerts from './hooks/useAlerts';
import InventoryViewAlertsDeleteAlert from './InventoryViewAlertsDeleteAlert';
import InventoryViewAlertDisplayAlerts from './InventoryViewAlertsDisplay';
import InventoryViewAlertsCreateOrEditAlert from './InventoryViewAlertsEditAlert';
import InventoryViewAlertsError from './InventoryViewAlertsError';
import InventoryViewAlertHasNoExistingAlerts from './InventoryViewAlertsHasNoAlerts';
import InventoryViewAlertsSkeleton from './InventoryViewAlertsSkeleton';
import InventoryViewAlertsChooseAlertMethod from './InventoryViewAlertsChooseAlertMethod';

type InventoryViewAlertsProps = {
  viewId: number;
  showToast: (toast: ToastProps) => void;
};

const viewControllerOptions = {
  NO_ALERTS_OR_EXSITING_ALERTS: 0,
  CHOOSE_ALERT_METHOD: 1,
  CREATE_OR_EDIT_ALERT: 2,
  DELETE_ALERT: 3
};

function InventoryViewAlerts({ viewId, showToast }: InventoryViewAlertsProps) {
  const {
    loading,
    error,
    hasAlerts,
    isSlackConfigured,
    alerts,
    alertsViewController,
    editAlert,
    alertMethod,
    currentAlert,
    setAlertMethodInAndIncrementViewController,
    setViewControllerToAlertsBaseView,
    setViewControllerToDeleteView,
    createOrEditAlert,
    incrementViewController,
    decrementViewController,
    closeAlert,
    fetchViewAlerts
  } = useAlerts({ viewId });

  if (loading) {
    return <InventoryViewAlertsSkeleton />;
  }
  if (error) {
    return <InventoryViewAlertsError fetchViewAlerts={fetchViewAlerts} />;
  }

  switch (alertsViewController) {
    case viewControllerOptions.NO_ALERTS_OR_EXSITING_ALERTS:
      if (!hasAlerts) {
        return (
          <InventoryViewAlertHasNoExistingAlerts
            incrementViewController={incrementViewController}
          />
        );
      }
      if (editAlert) {
        return (
          <InventoryViewAlertsCreateOrEditAlert
            alertMethod={alertMethod}
            setViewControllerOnSubmit={setViewControllerToAlertsBaseView}
            setViewControllerOnClickingBackButton={
              setViewControllerToAlertsBaseView
            }
            setViewControllerOnDelete={setViewControllerToDeleteView}
            currentAlert={currentAlert}
            closeAlert={closeAlert}
            viewId={viewId}
            showToast={showToast}
          />
        );
      }
      return (
        <InventoryViewAlertDisplayAlerts
          alerts={alerts}
          createOrEditAlert={createOrEditAlert}
          setViewControllerOnAddAlert={incrementViewController}
        />
      );

    case viewControllerOptions.CHOOSE_ALERT_METHOD:
      return (
        <InventoryViewAlertsChooseAlertMethod
          setAlertMethodInViewController={
            setAlertMethodInAndIncrementViewController
          }
          setViewControllerOnClickingBackButton={decrementViewController}
          isSlackConfigured={isSlackConfigured}
        />
      );
    case viewControllerOptions.CREATE_OR_EDIT_ALERT:
      return (
        <InventoryViewAlertsCreateOrEditAlert
          alertMethod={alertMethod}
          setViewControllerOnSubmit={setViewControllerToAlertsBaseView}
          setViewControllerOnClickingBackButton={decrementViewController}
          setViewControllerOnDelete={incrementViewController}
          currentAlert={currentAlert}
          closeAlert={closeAlert}
          viewId={viewId}
          showToast={showToast}
        />
      );
    case viewControllerOptions.DELETE_ALERT:
      return (
        <InventoryViewAlertsDeleteAlert
          alertMethod={alertMethod}
          viewControllerOnCancelButton={decrementViewController}
          currentAlert={currentAlert}
          closeAlert={closeAlert}
          viewId={viewId}
          showToast={showToast}
        />
      );
    default:
      return null;
  }
}

export default InventoryViewAlerts;
