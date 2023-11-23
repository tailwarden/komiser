import { capitalizeString } from '@utils/formatString';

export type SidepanelTabsProps = {
  goTo: (page: any) => void;
  page: string;
  tabs: string[];
};

function SidepanelTabs({ goTo, page, tabs }: SidepanelTabsProps) {
  return (
    <div className="border-b-2 border-gray-200 text-center text-sm font-medium text-gray-500">
      <ul className="-mb-[2px] flex justify-between sm:justify-start">
        {tabs.map((tab, idx) => (
          <li key={idx} className="mr-2">
            <a
              onClick={() => goTo(tab.toLowerCase())}
              className={`inline-block cursor-pointer select-none rounded-t-lg border-b-2 px-2 py-4 sm:p-4 
                     ${
                       page === tab.toLowerCase()
                         ? 'border-darkcyan-500 text-darkcyan-500 hover:text-darkcyan-500'
                         : 'border-transparent hover:text-darkcyan-700'
                     }`}
            >
              {capitalizeString(tab)} {/* capitalize first letter */}
            </a>
          </li>
        ))}
      </ul>
    </div>
  );
}

export default SidepanelTabs;
