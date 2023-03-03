import Button from '../../../../button/Button';

type InventoryViewAlertsErrorProps = {
  fetchViewAlerts: () => void;
};

function InventoryViewAlertsError({
  fetchViewAlerts
}: InventoryViewAlertsErrorProps) {
  return (
    <div className="flex flex-col gap-4">
      <div className="flex select-none items-center justify-between rounded-lg border border-black-170 p-6">
        <div className="flex items-center gap-4">
          <div className="flex flex-col gap-2">
            <p className="text-sm text-black-400">
              There was an error fetching the Slack alerts
            </p>
          </div>
        </div>
        <Button style="outline" size="sm" onClick={fetchViewAlerts}>
          <svg
            xmlns="http://www.w3.org/2000/svg"
            width="16"
            height="16"
            fill="none"
            viewBox="0 0 24 24"
          >
            <path
              stroke="currentColor"
              strokeLinecap="round"
              strokeLinejoin="round"
              strokeWidth="1.5"
              d="M22 12c0 5.52-4.48 10-10 10s-8.89-5.56-8.89-5.56m0 0h4.52m-4.52 0v5M2 12C2 6.48 6.44 2 12 2c6.67 0 10 5.56 10 5.56m0 0v-5m0 5h-4.44"
            ></path>
          </svg>
          Try again
        </Button>
      </div>
    </div>
  );
}

export default InventoryViewAlertsError;
