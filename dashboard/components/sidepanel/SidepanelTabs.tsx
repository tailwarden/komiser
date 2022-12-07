type SidepanelTabsProps = {
  goTo: (page: any) => void;
  page: string;
  tabs: string[];
};

function SidepanelTabs({ goTo, page, tabs }: SidepanelTabsProps) {
  return (
    <div className="text-sm font-medium text-center border-b-2 border-black-150 text-black-300">
      <ul className="flex justify-between sm:justify-start -mb-[2px]">
        {tabs.map((tab, idx) => (
          <li key={idx} className="mr-2">
            <a
              onClick={() => goTo(tab.toLowerCase())}
              className={`select-none inline-block py-4 px-2 sm:p-4 rounded-t-lg border-b-2 border-transparent hover:text-komiser-700 cursor-pointer 
                     ${
                       page === tab.toLowerCase() &&
                       `text-komiser-600 border-komiser-600 hover:text-komiser-600`
                     }`}
            >
              {tab}
            </a>
          </li>
        ))}
      </ul>
    </div>
  );
}

export default SidepanelTabs;
