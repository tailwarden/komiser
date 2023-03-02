import { View } from '../../../hooks/useInventory/types/useInventoryTypes';
import useSlackAlerts from './hooks/useSlackAlerts';
import InventoryViewAlertHasNoSlackAlerts from './InventoryViewAlertsHasNoSlackAlerts';
import InventoryViewAlertHasNoSlackIntegration from './InventoryViewAlertsHasNoSlackIntegration';

type InventoryViewAlertsProps = {
  view: View;
};

function InventoryViewAlerts({ view }: InventoryViewAlertsProps) {
  const { loading, error, hasSlack, slackAlerts, hasNoSlackAlerts } =
    useSlackAlerts({ view });

  if (loading) return <p>loading</p>;

  if (error) return <p>error fetching</p>;

  if (!hasSlack) return <InventoryViewAlertHasNoSlackIntegration />;

  if (hasNoSlackAlerts) return <InventoryViewAlertHasNoSlackAlerts />;

  return <p>Blep</p>;
}

export default InventoryViewAlerts;
