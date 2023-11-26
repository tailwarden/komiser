import { NextRouter } from 'next/router';
import Button from '../../button/Button';

type Props = {
  query: string;
  setQuery: (query: string) => void;
  router: NextRouter;
};

function InventorySearchNoResults({ query, setQuery, router }: Props) {
  return (
    <div className="flex w-full flex-wrap-reverse items-center justify-center gap-8 rounded-b-lg bg-white px-6 py-8 text-sm sm:py-20">
      <picture className="flex-shrink-0">
        <img
          src="/assets/img/purplin/reading.svg"
          className="w-32"
          alt="No results"
        />
      </picture>
      <div className="flex flex-col items-start gap-4">
        <div>
          {query ? (
            <>
              <p className="text-gray-500">No results were found for:</p>
              <p className="font-medium">{query}</p>
            </>
          ) : (
            <p className="text-gray-500">
              No results were found for this filter.
            </p>
          )}
        </div>
        <Button
          style="secondary"
          onClick={
            query
              ? () => {
                  setQuery('');
                }
              : () => router.push(router.pathname)
          }
        >
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
              strokeWidth="1.5"
              d="M7.757 16.243l8.486-8.486M16.243 16.243L7.757 7.757"
            ></path>
          </svg>
          {query ? 'Clear search' : 'Clear filters'}
        </Button>
      </div>
    </div>
  );
}

export default InventorySearchNoResults;
