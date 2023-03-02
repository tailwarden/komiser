import { SlackAlerts } from './hooks/useSlackAlerts';

type InventoryViewAlertsDisplayProps = {
  slackAlerts: SlackAlerts[] | undefined;
};

function InventoryViewAlertsDisplay({
  slackAlerts
}: InventoryViewAlertsDisplayProps) {
  return <div>InventoryViewAlertsDisplay</div>;
}

export default InventoryViewAlertsDisplay;
