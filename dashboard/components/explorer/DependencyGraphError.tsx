import Button from '@components/button/Button';

type DashboardDependencyGraphErrorProps = {
  fetch: () => void;
};

function DependencyGraphError({ fetch }: DashboardDependencyGraphErrorProps) {
  return (
    <>
      <div className={`w-full rounded-lg bg-white px-6 py-4 pb-6`}>
        <div className="-mx-6 flex items-center justify-between border-b border-gray-300 px-6 pb-4">
          <div>
            <p className="text-sm font-semibold text-gray-950">
              Dependency Graph
            </p>
            <div className="mt-1"></div>
            <p className="text-xs text-gray-500">
              Analyze account resource associations
            </p>
          </div>
          <div className="flex h-[60px] items-center"></div>
        </div>
        <div className="mt-8"></div>
        <div className="flex flex-col items-center justify-center">
          <svg
            className="h-20 w-20"
            fill="none"
            stroke="currentColor"
            strokeLinecap="round"
            strokeLinejoin="round"
            strokeWidth={1.5}
            viewBox="0 0 24 24"
          >
            <path stroke="none" d="M0 0h24v24H0z" />
            <path d="M4 8V6a2 2 0 012-2h2M4 16v2a2 2 0 002 2h2M16 4h2a2 2 0 012 2v2M16 20h2a2 2 0 002-2v-2M9 10h.01M15 10h.01M9.5 15.05a3.5 3.5 0 015 0" />
          </svg>
          <p className="text-sm font-semibold text-gray-950">
            Cannot fetch Relationships
          </p>
          <div className="m-2 flex-shrink-0">
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
        <div className="mt-12"></div>
      </div>
    </>
  );
}

export default DependencyGraphError;
