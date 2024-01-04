import SidepanelHeader from '@components/sidepanel/SidepanelHeader';
import SidepanelPage from '@components/sidepanel/SidepanelPage';
import Pill from '@components/pill/Pill';
import Button from '@components/button/Button';
import CloseIcon from '@components/icons/CloseIcon';
import PlusIcon from '@components/icons/PlusIcon';
import Sidepanel from '@components/sidepanel/Sidepanel';
import SidepanelTabs from '@components/sidepanel/SidepanelTabs';
import formatNumber from '@utils/formatNumber';
import {
  InventoryItem,
  Pages,
  Tag
} from '../hooks/useInventory/types/useInventoryTypes';
import InventoryTagWrapper from './InventoryTagWrapper';

type InventorySidePanelProps = {
  closeModal: () => void;
  data: InventoryItem | undefined;
  goTo: (page: Pages) => void;
  page: Pages;
  updateTags: (action?: 'delete') => void;
  tags: Tag[] | [] | undefined;
  handleChange: (newData: Partial<Tag>, id?: number) => void;
  removeTag: (id: number) => void;
  addNewTag: () => void;
  loading: boolean;
  deleteLoading: boolean;
  isOpen: boolean;
  bulkItems: [] | string[];
  updateBulkTags: (action?: 'delete' | undefined) => void;
  tabs: string[];
};

function InventorySidePanel({
  closeModal,
  data,
  goTo,
  page,
  updateTags,
  tags,
  handleChange,
  removeTag,
  addNewTag,
  loading,
  deleteLoading,
  isOpen,
  bulkItems,
  updateBulkTags,
  tabs
}: InventorySidePanelProps) {
  const getLastFetched = (date: string) => {
    const dateLastFetched = new Date(date);
    const today = new Date();
    const aMonthAgo = new Date(
      today.getFullYear(),
      today.getMonth() - 1,
      today.getDate()
    );
    const aWeekAgo = new Date(
      today.getFullYear(),
      today.getMonth(),
      today.getDate() - 7
    );
    let message;
    if (dateLastFetched > aMonthAgo) {
      message = 'Since last month';
    } else if (dateLastFetched > aWeekAgo) {
      message = 'Since last week';
    } else {
      message = 'More than a month ago';
    }
    return message;
  };

  return (
    <>
      <Sidepanel isOpen={isOpen} closeModal={closeModal}>
        {/* Modal headers */}
        {data && (
          <SidepanelHeader
            title={data.service}
            subtitle={data.name}
            closeModal={closeModal}
            href={data.link}
            cloudProvider={data.provider}
          >
            {!data && bulkItems && (
              <div className="flex flex-col gap-1 py-1">
                <p className="font-medium text-gray-950">
                  Managing tags for {formatNumber(bulkItems.length)}{' '}
                  {bulkItems.length > 1 ? 'resources' : 'resource'}
                </p>
                <p className="text-xs text-gray-500">
                  All actions will overwrite previous tags for these resources
                </p>
              </div>
            )}
          </SidepanelHeader>
        )}

        {/* Tabs */}
        <SidepanelTabs goTo={goTo} page={page} tabs={tabs} />

        {/* Tab Content */}
        {tabs.includes('resource details') && (
          <SidepanelPage page={page} param={'resource details'}>
            <div className="space-y-6 pt-1">
              <div className="space-y-2">
                <h2 className="font-['Noto Sans'] text-neutral-500 text-sm font-normal leading-tight">
                  Cloud account
                </h2>
                <h2 className="font-['Noto Sans'] text-neutral-900 text-sm font-normal leading-tight">
                  {!data && (
                    <p className="h-3 w-48 animate-pulse rounded-xl bg-cyan-200"></p>
                  )}
                  {data && <span>{data.account}</span>}
                </h2>
              </div>
              <div className="space-y-2">
                <h2 className="font-['Noto Sans'] text-neutral-500 text-sm font-normal leading-tight">
                  Region
                </h2>
                <h2 className="font-['Noto Sans'] text-neutral-900 text-sm font-normal leading-tight">
                  {!data && (
                    <p className="h-3 w-48 animate-pulse rounded-xl bg-cyan-200"></p>
                  )}
                  {data && <span>{data.region}</span>}
                </h2>
              </div>
              <div className="space-y-2">
                <h2 className="font-['Noto Sans'] text-neutral-500 text-sm font-normal leading-tight">
                  Cost
                </h2>
                <h2 className=" flex items-center gap-2 text-sm">
                  {!data && (
                    <p className="h-3 w-48 animate-pulse rounded-xl bg-cyan-200"></p>
                  )}
                  {data && <span>{data?.cost.toFixed(2)}$</span>}
                  {data && (
                    <Pill status="removed">
                      {getLastFetched(data.fetchedAt)}
                    </Pill>
                  )}
                </h2>
              </div>
              <div className="space-y-2">
                <h2 className="font-['Noto Sans'] text-neutral-500 text-sm font-normal leading-tight">
                  Relations
                </h2>
                <h2 className="font-['Noto Sans'] text-neutral-900 text-sm font-normal leading-tight">
                  {!data && (
                    <p className="h-3 w-48 animate-pulse rounded-xl bg-cyan-200"></p>
                  )}
                  {data && (
                    <span>{data.relations.length} related resources</span>
                  )}
                </h2>
              </div>
              {data && data.metadata !== null && (
              <div className="space-y-2">
                <h2 className="font-['Noto Sans'] text-neutral-500 text-sm font-normal leading-tight">
                  Metadata
                </h2>
                <h2 className="font-['Noto Sans'] text-neutral-900 text-sm font-normal leading-tight">
                  {!data && (
                    <p className="h-3 w-48 animate-pulse rounded-xl bg-cyan-200"></p>
                  )}
                  {data && (
                    <pre>
                      {JSON.stringify(data.metadata, null, 2)}
                    </pre>
                  )}
                </h2>
              </div>
              )}
            </div>
          </SidepanelPage>
        )}
        {/* Tags form */}
        {tabs.includes('tags') && (
          <SidepanelPage page={page} param={'tags'}>
            <form
              onSubmit={e => {
                e.preventDefault();

                if (!data && bulkItems) {
                  updateBulkTags();
                } else {
                  updateTags();
                }
              }}
              className="flex flex-col gap-6 px-1 pt-2"
            >
              {tags &&
                tags.map((tag, id) => (
                  <div key={id} className="flex items-center gap-4">
                    <InventoryTagWrapper
                      tag={tag}
                      id={id}
                      handleChange={handleChange}
                    />
                    {tags.length > 1 && (
                      <Button
                        size="xxs"
                        style="ghost"
                        onClick={() => removeTag(id)}
                      >
                        <CloseIcon width={24} height={24} />
                      </Button>
                    )}
                  </div>
                ))}
              <Button onClick={addNewTag} style="secondary" size="sm">
                <PlusIcon width={24} height={24} />
                Add new tag
              </Button>
              <div className="flex items-center justify-end gap-6 pt-2">
                {((data && data.tags && data.tags.length > 0) ||
                  (!data && bulkItems)) && (
                  <Button
                    size="lg"
                    style="delete"
                    loading={deleteLoading}
                    onClick={() => {
                      if (!data && bulkItems) {
                        goTo('delete');
                      } else {
                        updateTags('delete');
                      }
                    }}
                  >
                    {deleteLoading ? 'Deleting...' : 'Delete all tags'}
                  </Button>
                )}
                <Button
                  type="submit"
                  size="lg"
                  loading={loading}
                  disabled={
                    tags &&
                    !tags.every(tag => tag.key.trim() && tag.value.trim())
                  }
                >
                  {data && data.tags && data.tags.length > 0
                    ? 'Save changes'
                    : 'Add tags'}
                </Button>
              </div>
            </form>
          </SidepanelPage>
        )}
        <div>
          {page === 'delete' && (
            <>
              <div className="flex flex-col gap-6 bg-gray-50 p-6">
                <div className="mx-auto flex h-12 w-12 items-center justify-center rounded-full bg-red-50 text-red-500">
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
                      d="M12 22c5.5 0 10-4.5 10-10S17.5 2 12 2 2 6.5 2 12s4.5 10 10 10zM12 8v5M11.995 16h.009"
                    ></path>
                  </svg>
                </div>
                <div className="flex flex-col items-center gap-6">
                  <p className="text-center font-medium text-gray-950">
                    Are you sure you want to delete all tags from{' '}
                    {formatNumber(bulkItems.length)}{' '}
                    {bulkItems.length > 1 ? 'resources' : 'resource'}?
                  </p>
                  <p className="text-sm text-gray-700">
                    This is a permanent action, and it will also delete previous
                    tags you have added to these resources.
                  </p>
                </div>
                <div className="flex items-center justify-end gap-6">
                  <Button
                    size="lg"
                    style="ghost"
                    onClick={() => {
                      goTo('tags');
                    }}
                  >
                    Cancel
                  </Button>
                  <Button
                    size="lg"
                    style="delete"
                    loading={deleteLoading}
                    onClick={() => updateBulkTags('delete')}
                  >
                    {deleteLoading ? 'Deleting...' : 'Delete all tags'}
                  </Button>
                </div>
              </div>
            </>
          )}
        </div>
      </Sidepanel>
    </>
  );
}

export default InventorySidePanel;
