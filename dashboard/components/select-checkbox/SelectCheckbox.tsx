import { ChangeEvent, Dispatch, SetStateAction, useState } from 'react';
import Button from '../button/Button';
import Checkbox from '../checkbox/Checkbox';
import { ResourcesManagerQuery } from '../dashboard/components/resources-manager/hooks/useResourcesManager';
import useSelectCheckbox from './hooks/useSelectCheckbox';
import ChevronDownIcon from '../icons/ChevronDownIcon';

export type SelectCheckboxProps = {
  label: string;
  query: ResourcesManagerQuery;
  exclude: string[];
  setExclude: Dispatch<SetStateAction<string[]>>;
};

function SelectCheckbox({
  label,
  query,
  exclude,
  setExclude
}: SelectCheckboxProps) {
  const [isOpen, setIsOpen] = useState(false);
  const [search, setSearch] = useState('');
  const [checkedItems, setCheckedItems] = useState<string[]>([]);

  const { listOfExcludableItems, error } = useSelectCheckbox(query);

  function toggle() {
    setCheckedItems(exclude);
    setSearch('');
    setIsOpen(!isOpen);
  }

  function handleChange(e: ChangeEvent<HTMLInputElement>, item: string) {
    if (e.currentTarget.checked) {
      setCheckedItems(prev => [...prev, item]);
    } else {
      const findItem = checkedItems.find(resource => resource === item);
      const newCheckedItems = checkedItems.filter(
        resource => resource !== findItem
      );
      setCheckedItems(newCheckedItems);
    }
  }

  function handleCheckAll(e: ChangeEvent<HTMLInputElement>) {
    if (e.currentTarget.checked) {
      setCheckedItems(listOfExcludableItems);
    } else {
      setCheckedItems([]);
    }
  }

  function submit() {
    setExclude(checkedItems);
  }

  let resources = listOfExcludableItems;

  if (search) {
    resources = resources.filter(resource =>
      resource.toLowerCase().includes(search.toLowerCase())
    );
  }

  return (
    <div className="relative">
      <div
        className="pointer-events-none absolute bottom-[1.15rem]
        right-4 text-gray-950 transition-all"
      >
        <ChevronDownIcon width={24} height={24} />
      </div>
      <button
        onClick={toggle}
        className={`h-[60px] w-full overflow-hidden rounded text-left outline hover:outline-gray-300 focus:outline-2 focus:outline-darkcyan-500 ${
          isOpen ? 'outline-2 outline-darkcyan-500' : 'outline-gray-300'
        }`}
      >
        <div className="absolute right-0 top-1 h-[50px] w-6 bg-gradient-to-r from-transparent via-white to-white"></div>
        <span className="pointer-events-none absolute bottom-[1.925rem] left-4 origin-left scale-75 select-none font-normal text-gray-500">
          {label}
        </span>
        <div className="pointer-events-none flex w-full appearance-none items-center gap-2 rounded bg-white pb-[0.75rem] pl-4 pr-16 pt-[1.75rem] text-sm text-gray-950">
          {exclude.length > 0 &&
            exclude.slice(0, 3).map((resource, idx) => (
              <p
                key={idx}
                className="whitespace-nowrap rounded bg-gray-50 px-3 py-1 text-xs"
              >
                {resource}
              </p>
            ))}
          {exclude.length > 3 && (
            <p className="rounded-full bg-gray-50 px-3 py-1 text-xs">
              +{exclude.length - 3}
            </p>
          )}
          {exclude.length === 0 && 'None excluded'}
        </div>
      </button>

      {isOpen && (
        <>
          <div
            data-testid="overlay"
            onClick={toggle}
            className="fixed inset-0 z-20 hidden animate-fade-in bg-transparent opacity-0 sm:block"
          ></div>
          <div className="absolute top-[4.15rem] z-[21] w-full rounded-lg border border-gray-50 bg-white shadow-right">
            <div className="relative m-4 ">
              {!search ? (
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  width="24"
                  height="24"
                  fill="none"
                  viewBox="0 0 24 24"
                  className="absolute left-3 top-[0.5rem]"
                >
                  <path
                    d="M17 17L21 21"
                    stroke="#ababab"
                    strokeWidth="2"
                    strokeLinecap="round"
                    strokeLinejoin="round"
                  />
                  <path
                    d="M19 11C19 15.4183 15.4183 19 11 19C6.58172 19 3 15.4183 3 11C3 6.58172 6.58172 3 11 3C15.4183 3 19 6.58172 19 11Z"
                    stroke="#ababab"
                    strokeWidth="2"
                  />
                </svg>
              ) : (
                <svg
                  onClick={() => setSearch('')}
                  xmlns="http://www.w3.org/2000/svg"
                  width="24"
                  height="24"
                  fill="none"
                  viewBox="0 0 24 24"
                  className="absolute left-3 top-[0.5rem] cursor-pointer"
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
                value={search}
                onChange={e => setSearch(e.target.value)}
                type="text"
                placeholder="Search"
                className="h-10 w-full rounded-md border border-gray-300 bg-white py-4 pl-10 pr-6 text-sm text-gray-950 caret-darkcyan-700 placeholder:text-gray-500 focus:outline-none"
                autoFocus
                maxLength={64}
              />
            </div>
            {error && (
              <p className="text-sm text-gray-700">
                There was an error fetching the options for: {query}
              </p>
            )}
            {!error && (
              <>
                {!search && (
                  <div className="m-4 ml-6 flex items-center gap-2 text-sm">
                    <Checkbox
                      id="all"
                      onChange={e => {
                        handleCheckAll(e);
                      }}
                      checked={checkedItems.length === resources.length}
                    />
                    <label
                      htmlFor="all"
                      className="w-full text-sm text-gray-700"
                    >
                      Exclude All
                    </label>
                  </div>
                )}
                <hr className="bg-neutral-100 m-4 mb-0 h-px border-t-0 opacity-100 dark:opacity-50" />
                <div className="scrollbar mb-2 mr-3 mt-2 overflow-auto">
                  <div className="mt-2 flex max-h-[12rem] flex-col gap-3 p-4 pb-4 pl-6 pt-0">
                    {resources.map((resource, idx) => (
                      <div
                        key={idx}
                        className="flex items-center gap-2 text-sm"
                      >
                        <Checkbox
                          id={resource}
                          onChange={e => handleChange(e, resource)}
                          checked={
                            !!checkedItems.find(value => value === resource)
                          }
                        />
                        <label
                          htmlFor={resource}
                          className="w-full text-gray-700"
                        >
                          {resource}
                        </label>
                      </div>
                    ))}
                    {resources.length === 0 && (
                      <p className="text-sm text-gray-700">
                        There are no results for {search}
                      </p>
                    )}
                  </div>
                </div>
                <div className="flex flex-col border-t border-gray-300 p-4">
                  <Button onClick={submit}>Apply</Button>
                </div>
              </>
            )}
          </div>
        </>
      )}
    </div>
  );
}

export default SelectCheckbox;
