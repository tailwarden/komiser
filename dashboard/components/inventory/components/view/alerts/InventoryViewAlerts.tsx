import useSlackAlerts from './hooks/useSlackAlerts';
import InventoryViewAlertsDisplay from './InventoryViewAlertsDisplay';
import InventoryViewAlertsEditSlackAlert from './InventoryViewAlertsEditSlackAlert';
import InventoryViewAlertHasNoSlackAlerts from './InventoryViewAlertsHasNoSlackAlerts';
import InventoryViewAlertHasNoSlackIntegration from './InventoryViewAlertsHasNoSlackIntegration';
import InventoryViewAlertsSkeleton from './InventoryViewAlertsSkeleton';

type InventoryViewAlertsProps = {
  viewId: number;
};

function InventoryViewAlerts({ viewId }: InventoryViewAlertsProps) {
  const {
    loading,
    error,
    hasSlack,
    slackAlerts,
    hasNoSlackAlerts,
    editSlackAlert,
    currentSlackAlert,
    createOrEditSlackAlert,
    closeSlackAlert
  } = useSlackAlerts({ viewId });

  if (loading) return <InventoryViewAlertsSkeleton />;

  if (error) return <p>error fetching</p>;

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
