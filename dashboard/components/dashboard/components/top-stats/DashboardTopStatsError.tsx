import Button from '../../../button/Button';

type DashboardTopStatsErrorProps = {
  fetch: () => void;
};

function DashboardTopStatsError({ fetch }: DashboardTopStatsErrorProps) {
  return (
    <div
      data-testid="error"
      className="flex h-[7.5rem] items-center justify-center gap-4 rounded-lg bg-white"
    >
      <p className="text-sm text-gray-700">
        There was an error loading the top stats.
      </p>
      <div className="flex-shrink-0">
        <Button style="secondary" size="sm" onClick={fetch}>
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

export default DashboardTopStatsError;
