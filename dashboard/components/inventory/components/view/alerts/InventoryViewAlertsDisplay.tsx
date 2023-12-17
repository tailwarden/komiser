import Image from 'next/image';
import formatNumber from '../../../../../utils/formatNumber';
import Button from '../../../../button/Button';
import ChevronRightIcon from '../../../../icons/ChevronRightIcon';
import { Alert } from './hooks/useAlerts';

type InventoryViewAlertDisplayAlertsProps = {
  alerts: Alert[] | undefined;
  createOrEditAlert: (alertId?: number | undefined) => void;
  setViewControllerOnAddAlert: () => void;
};

function InventoryViewAlertDisplayAlerts({
  alerts,
  createOrEditAlert,
  setViewControllerOnAddAlert
}: InventoryViewAlertDisplayAlertsProps) {
  return (
    <div className="flex flex-col gap-4">
      {alerts?.map(alert => (
        <div
          onClick={() => createOrEditAlert(alert.id)}
          key={alert.id}
          className="flex cursor-pointer select-none items-center justify-between rounded-lg border border-gray-200 p-6 hover:border-gray-300"
        >
          <div className="flex items-center gap-4">
            <Image
              src={
                alert.isSlack
                  ? '/assets/img/others/slack.svg'
                  : '/assets/img/others/custom-webhook.svg'
              }
              height={42}
              width={42}
              alt={alert.isSlack ? 'Slack logo' : 'Webhook logo'}
            />
            <div className="flex flex-col">
              <p className="font-semibold text-gray-950">{alert.name}</p>
              <p className="text-xs text-gray-700">
                {alert.budget
                  ? `When total cost is over $${formatNumber(
                      Number(alert.budget)
                    )}`
                  : `When cloud resources are over ${formatNumber(
                      Number(alert.usage)
                    )}`}
              </p>
            </div>
          </div>
          <ChevronRightIcon width={24} height={24} />
        </div>
      ))}
      <div className="self-end">
        <Button size="lg" onClick={setViewControllerOnAddAlert}>
          Add alert
        </Button>
      </div>
    </div>
  );
}

export default InventoryViewAlertDisplayAlerts;
