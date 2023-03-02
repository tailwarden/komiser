import { View } from '../../../hooks/useInventory/types/useInventoryTypes';
import useSlackAlerts from './hooks/useSlackAlerts';
import InventoryViewAlertsDisplay from './InventoryViewAlertsDisplay';
import InventoryViewAlertsEditSlackAlert from './InventoryViewAlertsEditSlackAlert';
import InventoryViewAlertHasNoSlackAlerts from './InventoryViewAlertsHasNoSlackAlerts';
import InventoryViewAlertHasNoSlackIntegration from './InventoryViewAlertsHasNoSlackIntegration';

type InventoryViewAlertsProps = {
  view: View;
};

function InventoryViewAlerts({ view }: InventoryViewAlertsProps) {
  const {
    loading,
    error,
    hasSlack,
    slackAlerts,
    hasNoSlackAlerts,
    editSlackAlert,
    createSlackAlert,
    closeSlackAlert
  } = useSlackAlerts({ view });

  if (loading) return <p>loading</p>;

  if (error) return <p>error fetching</p>;

  if (!hasSlack) return <InventoryViewAlertHasNoSlackIntegration />;

  if (hasNoSlackAlerts)
    return (
      <InventoryViewAlertHasNoSlackAlerts createSlackAlert={createSlackAlert} />
    );

  if (editSlackAlert)
    return (
      <InventoryViewAlertsEditSlackAlert
        closeSlackAlert={closeSlackAlert}
        viewId={view.id}
      />
    );

  return <InventoryViewAlertsDisplay slackAlerts={slackAlerts} />;
}

export default InventoryViewAlerts;
