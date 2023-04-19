import ExportCSV from '../../export-csv/ExportCSV';
import CloseIcon from '../../icons/CloseIcon';
import SearchIcon from '../../icons/SearchIcon';
import { ToastProps } from '../../toast/hooks/useToast';

type InventorySearchBarProps = {
  query: string;
  setQuery: (query: string) => void;
  error: boolean;
  setToast: (toast: ToastProps | undefined) => void;
};

function InventorySearchBar({
  query,
  setQuery,
  error,
  setToast
}: InventorySearchBarProps) {
  return (
    <>
      {!error && (
        <div className="relative rounded-lg rounded-b-none">
          {!query ? (
            <div className="absolute top-[1.625rem] left-6">
              <SearchIcon width={16} height={16} />
            </div>
          ) : (
            <div
              className="absolute top-[1.7rem] left-6 cursor-pointer"
              onClick={() => setQuery('')}
            >
              <CloseIcon width={16} height={16} />
            </div>
          )}

          <input
            value={query}
            onChange={e => setQuery(e.target.value)}
            type="text"
            placeholder="Search by tags, service, name, region..."
            className="w-full rounded-t-lg border-b border-black-200/30 bg-white py-6 pl-14 pr-6 text-sm text-black-900 caret-secondary placeholder:text-black-300 focus:outline-none"
            autoComplete="off"
            data-lpignore="true"
            data-form-type="other"
          />
          <div className="absolute top-[14px] right-4">
            <ExportCSV setToast={setToast} />
          </div>
        </div>
      )}
    </>
  );
}

export default InventorySearchBar;
