type InventorySearchBarProps = {
  query: string;
  setQuery: (query: string) => void;
  error: boolean;
};

function InventorySearchBar({
  query,
  setQuery,
  error
}: InventorySearchBarProps) {
  return (
    <>
      {!error && (
        <div className="relative overflow-hidden rounded-lg rounded-b-none">
          {!query ? (
            <svg
              xmlns="http://www.w3.org/2000/svg"
              width="16"
              height="16"
              fill="none"
              viewBox="0 0 24 24"
              className="absolute top-[1.125rem] left-6"
            >
              <path
                stroke="currentColor"
                strokeLinecap="round"
                strokeLinejoin="round"
                strokeWidth="1.5"
                d="M11.5 21a9.5 9.5 0 100-19 9.5 9.5 0 000 19zM22 22l-2-2"
              ></path>
            </svg>
          ) : (
            <svg
              onClick={() => setQuery('')}
              xmlns="http://www.w3.org/2000/svg"
              width="16"
              height="16"
              fill="none"
              viewBox="0 0 24 24"
              className="absolute top-[1.175rem] left-6 cursor-pointer"
            >
              <path
                stroke="currentColor"
                strokeLinecap="round"
                strokeLinejoin="round"
                strokeWidth="1.5"
                d="M7.757 16.243l8.486-8.486M16.243 16.243L7.757 7.757"
              ></path>
            </svg>
          )}

          <input
            value={query}
            onChange={e => setQuery(e.target.value)}
            type="text"
            placeholder="Search by tags, service, name, region..."
            className="w-full border-b border-black-200/30 bg-white py-4 pl-14 pr-6 text-sm text-black-900 caret-secondary placeholder:text-black-300 focus:outline-none"
          />
        </div>
      )}
    </>
  );
}

export default InventorySearchBar;
