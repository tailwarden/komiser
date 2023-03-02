import Image from 'next/image';
import { SlackAlerts } from './hooks/useSlackAlerts';

type InventoryViewAlertsDisplayProps = {
  slackAlerts: SlackAlerts[] | undefined;
};

function InventoryViewAlertsDisplay({
  slackAlerts
}: InventoryViewAlertsDisplayProps) {
  return (
    <>
      {slackAlerts?.map(alert => (
        <div
          key={alert.id}
          className="flex cursor-pointer select-none items-center justify-between rounded-lg border border-black-170 p-6 hover:border-black-200"
        >
          <div className="flex items-center gap-4">
            <Image
              src="./assets/img/others/slack.svg"
              height={42}
              width={42}
              alt="Slack logo"
            />
            <div className="flex flex-col">
              <p className="font-semibold text-black-900">{alert.name}</p>
              <p className="text-xs text-black-400">
                {alert.budget
                  ? `When total cost is over $${alert.budget}`
                  : `When cloud resources are over ${alert.usage}`}
              </p>
            </div>
          </div>
          <svg
            xmlns="http://www.w3.org/2000/svg"
            width="24"
            height="24"
            fill="none"
            viewBox="0 0 24 24"
          >
            <path
              stroke="currentColor"
              strokeLinecap="round"
              strokeLinejoin="round"
              strokeMiterlimit="10"
              strokeWidth="1.5"
              d="M8.91 19.92l6.52-6.52c.77-.77.77-2.03 0-2.8L8.91 4.08"
            ></path>
          </svg>
        </div>
      ))}
    </>
  );
}

export default InventoryViewAlertsDisplay;
