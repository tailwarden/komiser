import { NextRouter } from 'next/router';
import { ViewProps } from '../hooks/useInventory';

type InventoryViewsTabsProps = {
  views: ViewProps[] | undefined;
  router: NextRouter;
};

function InventoryViewsTabs({ views, router }: InventoryViewsTabsProps) {
  return (
    <>
      <div className="text-center text-sm font-medium text-black-300">
        <ul className="-mb-[2px] flex flex-wrap justify-between sm:justify-start">
          <li>
            <a
              onClick={() => {
                if (router.asPath !== '/') router.push('/');
              }}
              className={`inline-block cursor-pointer select-none rounded-t-lg border-b-2 border-transparent py-4 px-2 hover:text-komiser-700 sm:p-4 
                       ${
                         !router.query.view &&
                         `border-komiser-600 text-komiser-600 hover:text-komiser-600`
                       }`}
            >
              All resources
            </a>
          </li>
          {views &&
            views.length > 0 &&
            views.map((view, idx) => (
              <li key={idx}>
                <a
                  onClick={() => {
                    if (router.query.view !== view.id.toString()) {
                      router.push(`/?view=${view.id}`);
                    }
                  }}
                  className={`inline-block cursor-pointer select-none whitespace-nowrap rounded-t-lg border-b-2 border-transparent py-4 px-2 hover:text-komiser-700 sm:p-4
                       ${
                         router.query.view === view.id.toString() &&
                         `border-komiser-600 text-komiser-600 hover:text-komiser-600`
                       }`}
                >
                  {view.name}
                </a>
              </li>
            ))}
        </ul>
      </div>
    </>
  );
}

export default InventoryViewsTabs;
