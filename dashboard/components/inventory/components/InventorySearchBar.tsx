import { ToastProps } from '@components/toast/Toast';
import ExportCSV from '../../export-csv/ExportCSV';
import CloseIcon from '../../icons/CloseIcon';
import SearchIcon from '../../icons/SearchIcon';

type InventorySearchBarProps = {
  query: string;
  setQuery: (query: string) => void;
  error: boolean;
  showToast: (toast: ToastProps) => void;
};

function InventorySearchBar({
  query,
  setQuery,
  error,
  showToast
}: InventorySearchBarProps) {
  return (
    <>
      {!error && (
        <div className="relative rounded-lg rounded-b-none">
          {!query ? (
            <div className="absolute left-6 top-[1.625rem]">
              <SearchIcon width={16} height={16} />
            </div>
          ) : (
            <div
              className="absolute left-6 top-[1.7rem] cursor-pointer"
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
            className="w-full rounded-t-lg border-b border-gray-300 bg-white py-6 pl-14 pr-6 text-sm text-gray-950 caret-darkcyan-700 placeholder:text-gray-500 focus:outline-none"
            autoComplete="off"
            data-lpignore="true"
            data-form-type="other"
            maxLength={64}
          />
          <div className="absolute right-4 top-[14px]">
            <ExportCSV showToast={showToast} />
          </div>
        </div>
      )}
    </>
  );
}

export default InventorySearchBar;
