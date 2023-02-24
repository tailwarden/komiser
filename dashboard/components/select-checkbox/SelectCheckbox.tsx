import { ChangeEvent, Dispatch, SetStateAction, useState } from 'react';
import Button from '../button/Button';
import Checkbox from '../checkbox/Checkbox';
import { ResourcesManagerQuery } from '../dashboard/components/resources-manager/hooks/useResourcesManager';
import useSelectCheckbox from './hooks/useSelectCheckbox';

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
      <button
        onClick={toggle}
        className={`h-[60px] w-full overflow-hidden rounded text-left outline hover:outline-black-200 focus:outline-2 focus:outline-primary ${
          isOpen ? 'outline-2 outline-primary' : 'outline-black-200/50'
        }`}
      >
        <div className="absolute right-0 top-1 h-[50px] w-6 bg-gradient-to-r from-transparent via-white to-white"></div>
        <span className="pointer-events-none absolute left-4 bottom-[1.925rem] origin-left scale-75 select-none font-normal text-black-300">
          {label}
        </span>
        <div className="pointer-events-none flex w-full appearance-none items-center gap-2 rounded bg-white pt-[1.75rem] pb-[0.75rem] pl-4 pr-16 text-sm text-black-900">
          {exclude.length > 0 &&
            exclude.slice(0, 3).map((resource, idx) => (
              <p
                key={idx}
                className="whitespace-nowrap rounded bg-black-100 py-1 px-3 text-xs"
              >
                {resource}
              </p>
            ))}
          {exclude.length > 3 && (
            <p className="rounded-full bg-black-100 py-1 px-3 text-xs">
              +{exclude.length - 3}
            </p>
          )}
          {exclude.length === 0 && 'None excluded'}
        </div>
      </button>

      {isOpen && (
        <>
          <div
            onClick={toggle}
            className="fixed inset-0 z-20 hidden animate-fade-in bg-transparent opacity-0 sm:block"
          ></div>
          <div className="absolute top-[4.15rem] z-[21] w-full rounded-lg border border-black-200 bg-white shadow-lg">
            <div className="relative overflow-hidden rounded-lg rounded-b-none">
              {!search ? (
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  width="16"
                  height="16"
                  fill="none"
                  viewBox="0 0 24 24"
                  className="absolute top-[1.125rem] left-4"
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
                  onClick={() => setSearch('')}
                  xmlns="http://www.w3.org/2000/svg"
                  width="16"
                  height="16"
                  fill="none"
                  viewBox="0 0 24 24"
                  className="absolute top-[1.175rem] left-4 cursor-pointer"
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
                className="w-full border-b border-black-200/50 bg-white py-4 pl-10 pr-6 text-sm text-black-900 caret-secondary placeholder:text-black-300 focus:outline-none"
                autoFocus
              />
            </div>
            {error && (
              <p className="text-sm text-black-400">
                There was an error fetching the options for: {query}
              </p>
            )}
            {!error && (
              <>
                <div className="flex max-h-[12rem] flex-col gap-3 overflow-auto p-4">
                  {resources.map((resource, idx) => (
                    <div key={idx} className="flex items-center gap-2 text-sm">
                      <Checkbox
                        id={resource}
                        onChange={e => handleChange(e, resource)}
                        checked={
                          !!checkedItems.find(value => value === resource)
                        }
                      />
                      <label htmlFor={resource} className="w-full">
                        {resource}
                      </label>
                    </div>
                  ))}
                  {resources.length === 0 && (
                    <p className="text-sm text-black-400">
                      There are no results for {search}
                    </p>
                  )}
                </div>
                <div className="flex flex-col border-t border-black-200/50 p-4">
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
