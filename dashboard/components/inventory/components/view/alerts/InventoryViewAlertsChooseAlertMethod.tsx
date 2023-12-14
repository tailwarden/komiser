import Image from 'next/image';
import ChevronRightIcon from '../../../../icons/ChevronRightIcon';
import ArrowLeftIcon from '../../../../icons/ArrowLeftIcon';
import { AlertMethod } from './hooks/useAlerts';

type InventoryViewAlertsChooseAlertMethodProps = {
  setAlertMethodInViewController: (alertName: AlertMethod) => void;
  setViewControllerOnClickingBackButton: () => void;
  isSlackConfigured: boolean;
};

function InventoryViewAlertsChooseAlertMethod({
  setAlertMethodInViewController,
  setViewControllerOnClickingBackButton,
  isSlackConfigured
}: InventoryViewAlertsChooseAlertMethodProps) {
  const webhookOptions = [
    {
      id: AlertMethod.SLACK,
      name: 'Slack',
      message: 'Get directly notified to take action',
      image: '/assets/img/others/slack.svg',
      alt: 'Slack logo'
    },
    {
      id: AlertMethod.WEBHOOK,
      name: 'Webhook',
      message: 'Integrate actions into your system',
      image: '/assets/img/others/custom-webhook.svg',
      alt: 'Webhook logo'
    }
  ];

  return (
    <div className="flex flex-col gap-4">
      <div
        onClick={() => setViewControllerOnClickingBackButton()}
        className="flex cursor-pointer items-center gap-2 self-start text-sm text-gray-950"
      >
        <ArrowLeftIcon width={24} height={24} />
        Pick a Handler
      </div>

      {webhookOptions?.map(alert => (
        <div key={alert.id}>
          <div
            onClick={() => {
              if (alert.id !== AlertMethod.SLACK || isSlackConfigured) {
                setAlertMethodInViewController(alert.id);
              }
            }}
            className={`flex cursor-pointer select-none items-center justify-between rounded-lg border border-gray-200 p-6 hover:border-gray-300 
                        ${
                          alert.id === AlertMethod.SLACK && !isSlackConfigured
                            ? 'pointer-events-none bg-gray-200'
                            : ''
                        }`}
          >
            <div className="flex items-center gap-4">
              <Image src={alert.image} height={42} width={42} alt={alert.alt} />
              <div className="flex flex-col">
                <p className="font-semibold text-gray-950">{alert.name}</p>
                <p className="text-xs text-gray-700">{alert.message}</p>
              </div>
            </div>
            <ChevronRightIcon width={24} height={24} />
          </div>

          {alert.id === AlertMethod.SLACK && !isSlackConfigured && (
            <div className="mt-2">
              <p className="text-xs text-gray-700">
                You have not set up your Slack integration. Learn how through
                our{' '}
                <a
                  href="https://docs.komiser.io/guides/alerts#slack-integration"
                  target="_blank"
                  rel="noreferrer"
                  className="text-darkcyan-500"
                >
                  <u>guide</u>
                </a>
                .
              </p>
            </div>
          )}
        </div>
      ))}
    </div>
  );
}

export default InventoryViewAlertsChooseAlertMethod;
