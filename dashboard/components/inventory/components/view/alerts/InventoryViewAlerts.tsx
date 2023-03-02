import useSlackAlerts from './hooks/useSlackAlerts';
import InventoryViewAlertHasNoSlack from './InventoryViewAlertsHasNoSlack';

function InventoryViewAlerts() {
  const { loading, error, hasSlack } = useSlackAlerts();

  if (loading) return <p>loading</p>;

  if (error) return <p>error fetching</p>;

  if (!hasSlack) return <InventoryViewAlertHasNoSlack />;

  return <p>Blep</p>;
}

export default InventoryViewAlerts;
