import { Tag } from '../hooks/useInventory';

type InventoryTableTagsProps = {
  tags: [] | Tag[] | null;
  setQuery: (query: string) => void;
};

function InventoryTableTags({ tags, setQuery }: InventoryTableTagsProps) {
  return (
    <>
      {tags && tags.length > 0 && (
        <div className="relative group">
          <div className="relative flex items-center gap-1 py-4 px-6">
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
                strokeWidth="2"
                d="M4.17 15.3l4.53 4.53a4.78 4.78 0 006.75 0l4.39-4.39a4.78 4.78 0 000-6.75L15.3 4.17a4.75 4.75 0 00-3.6-1.39l-5 .24c-2 .09-3.59 1.68-3.69 3.67l-.24 5c-.06 1.35.45 2.66 1.4 3.61z"
              ></path>
              <path
                stroke="currentColor"
                strokeLinecap="round"
                strokeWidth="2"
                d="M9.5 12a2.5 2.5 0 100-5 2.5 2.5 0 000 5z"
              ></path>
            </svg>
            <span className="absolute top-3 left-[2.375rem] flex items-center justify-center bg-komiser-200 rounded-full h-4 w-4 text-[10px] text-black-900 font-bold">
              {tags.length}
            </span>
          </div>
          <div className="absolute hidden group-hover:flex flex-col gap-2 rounded-lg right-6 top-11 z-10 bg-black-900 py-3 px-4 shadow-lg">
            {tags.map((tag, index) => (
              <div
                key={index}
                className="flex items-center gap-2 text-xs border-t border-white/20 -mx-4 px-4 pt-2  text-black-200 first:pt-0 first:border-none"
              >
                <div className="flex items-center gap-1">
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    width="14"
                    height="14"
                    fill="none"
                    viewBox="0 0 24 24"
                    className="text-black-400"
                  >
                    <path
                      stroke="currentColor"
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      strokeWidth="2"
                      d="M4.17 15.3l4.53 4.53a4.78 4.78 0 006.75 0l4.39-4.39a4.78 4.78 0 000-6.75L15.3 4.17a4.75 4.75 0 00-3.6-1.39l-5 .24c-2 .09-3.59 1.68-3.69 3.67l-.24 5c-.06 1.35.45 2.66 1.4 3.61z"
                    ></path>
                    <path
                      stroke="currentColor"
                      strokeLinecap="round"
                      strokeWidth="2"
                      d="M9.5 12a2.5 2.5 0 100-5 2.5 2.5 0 000 5z"
                    ></path>
                  </svg>
                  <span
                    onClick={e => {
                      setQuery(tag.key);
                    }}
                    className="hover:text-komiser-500 cursor-pointer"
                  >
                    {tag.key}:
                  </span>
                </div>
                <span
                  onClick={() => setQuery(tag.value)}
                  className="font-medium hover:text-komiser-500 cursor-pointer"
                >
                  {tag.value}
                </span>
              </div>
            ))}
          </div>
        </div>
      )}
    </>
  );
}

export default InventoryTableTags;
