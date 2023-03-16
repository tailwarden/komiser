import type { ToastProps } from '@/components/toast/hooks/useToast';
import useSlackAlerts from './hooks/useSlackAlerts';
import InventoryViewAlertsDisplay from './InventoryViewAlertsDisplay';
import InventoryViewAlertsEditSlackAlert from './InventoryViewAlertsEditSlackAlert';
import InventoryViewAlertsError from './InventoryViewAlertsError';
import InventoryViewAlertHasNoSlackAlerts from './InventoryViewAlertsHasNoSlackAlerts';
import InventoryViewAlertHasNoSlackIntegration from './InventoryViewAlertsHasNoSlackIntegration';
import InventoryViewAlertsSkeleton from './InventoryViewAlertsSkeleton';

type InventoryViewAlertsProps = {
  viewId: number;
  setToast: (toast: ToastProps | undefined) => void;
};

function InventoryViewAlerts({ viewId, setToast }: InventoryViewAlertsProps) {
  const {
    loading,
    error,
    hasSlack,
    slackAlerts,
    hasNoSlackAlerts,
    editSlackAlert,
    currentSlackAlert,
    createOrEditSlackAlert,
    closeSlackAlert,
    fetchViewAlerts
  } = useSlackAlerts({ viewId });

  if (loading) return <InventoryViewAlertsSkeleton />;

  if (error)
    return <InventoryViewAlertsError fetchViewAlerts={fetchViewAlerts} />;

  if (!hasSlack) return <InventoryViewAlertHasNoSlackIntegration />;

  if (hasNoSlackAlerts)
    return (
      <InventoryViewAlertHasNoSlackAlerts
        createOrEditSlackAlert={createOrEditSlackAlert}
      />
    );

  if (editSlackAlert)
    return (
      <InventoryViewAlertsEditSlackAlert
        currentSlackAlert={currentSlackAlert}
        closeSlackAlert={closeSlackAlert}
        viewId={viewId}
        setToast={setToast}
      />
    );

  return (
    <InventoryViewAlertsDisplay
      slackAlerts={slackAlerts}
      createOrEditSlackAlert={createOrEditSlackAlert}
    />
  );
}

export default InventoryViewAlerts;
