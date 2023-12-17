import Button from '../../../button/Button';

type DashboardCloudMapErrorProps = {
  fetch: () => void;
};

function DashboardCloudMapError({ fetch }: DashboardCloudMapErrorProps) {
  return (
    <div
      data-testid="error"
      className="min-h-[396px] w-full rounded-lg bg-white px-6 py-4 pb-6"
    >
      <div className="-mx-6 flex items-center justify-center gap-6 border-b border-gray-300 px-6 pb-4">
        <p className="text-sm text-gray-700">
          There was an error loading the cloud map.
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
        <div className="h-[60px]"></div>
      </div>
      <div className="mt-8"></div>
      <div className="-mx-6 -ml-20 min-w-full">
        <picture>
          <img src="/assets/img/others/world.svg" alt="" />
        </picture>
      </div>
      <div className="mt-12"></div>
      <div className="flex gap-4 text-xs text-gray-500">
        <div className="flex items-center gap-2">
          <div className="h-2 w-2 rounded-full bg-blue-500"></div>Active region
        </div>
        <div className="flex items-center gap-2">
          <div className="h-2 w-2 rounded-full bg-gray-500"></div>
          Inactive region
        </div>
      </div>
    </div>
  );
}

export default DashboardCloudMapError;
