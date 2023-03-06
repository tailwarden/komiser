import formatNumber from '../../../utils/formatNumber/formatNumber';
import providers from '../../../utils/providerHelper';
import Button from '../../button/Button';
import Sidepanel from '../../sidepanel/Sidepanel';
import SidepanelTabs from '../../sidepanel/SidepanelTabs';
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
  updateBulkTags
}: InventorySidePanelProps) {
  return (
    <>
      <Sidepanel isOpen={isOpen} closeModal={closeModal}>
        {/* Modal headers */}
        <div className="flex flex-wrap-reverse items-center justify-between gap-6 sm:flex-nowrap">
          {data && (
            <div className="flex flex-wrap items-center gap-4 sm:flex-nowrap">
              <picture className="flex-shrink-0">
                <img
                  src={providers.providerImg(data.provider)}
                  className="h-8 w-8 rounded-full"
                  alt={data.provider}
                />
              </picture>

              <div className="flex flex-col gap-1 py-1">
                <p className="... w-48 truncate font-medium text-black-900">
                  {data.service}
                </p>
                <p className="flex items-center gap-2 text-xs text-black-300">
                  {data.name}
                  <a
                    target="_blank"
                    href={data.link}
                    rel="noreferrer"
                    className="hover:text-primary"
                  >
                    <svg
                      xmlns="http://www.w3.org/2000/svg"
                      width="16"
                      height="16"
                      fill="none"
                      viewBox="0 0 24 24"
                    >
                      <path
                        stroke="currentColor"
                        strokeLinecap="round"
                        strokeLinejoin="round"
                        strokeWidth="1.5"
                        d="M13 11l8.2-8.2M22 6.8V2h-4.8M11 2H9C4 2 2 4 2 9v6c0 5 2 7 7 7h6c5 0 7-2 7-7v-2"
                      ></path>
                    </svg>
                  </a>
                </p>
              </div>
            </div>
          )}
          {!data && bulkItems && (
            <div className="flex flex-col gap-1 py-1">
              <p className="font-medium text-black-900">
                Managing tags for {formatNumber(bulkItems.length)}{' '}
                {bulkItems.length > 1 ? 'resources' : 'resource'}
              </p>
              <p className="text-xs text-black-300">
                All actions will overwrite previous tags for these resources
              </p>
            </div>
          )}

          <div className="flex flex-shrink-0 items-center gap-2">
            <Button style="outline" onClick={closeModal}>
              Close
            </Button>
          </div>
        </div>

        {/* Tabs */}
        <SidepanelTabs goTo={goTo} page={page} tabs={['Tags']} />

        {/* Tags form */}
        <div className="rounded-lg bg-black-100 p-6">
          <div className="flex flex-col gap-6">
            {page === 'tags' && (
              <form
                onSubmit={e => {
                  e.preventDefault();

                  if (!data && bulkItems) {
                    updateBulkTags();
                  } else {
                    updateTags();
                  }
                }}
                className="flex flex-col gap-6"
              >
                {tags &&
                  tags.map((tag, id) => (
                    <div key={id} className="flex items-center gap-6">
                      <InventoryTagWrapper
                        tag={tag}
                        id={id}
                        handleChange={handleChange}
                      />
                      {tags.length > 1 && (
                        <Button
                          size="xs"
                          style="ghost"
                          onClick={() => removeTag(id)}
                        >
                          <svg
                            xmlns="http://www.w3.org/2000/svg"
                            width="20"
                            height="20"
                            fill="none"
                            viewBox="0 0 24 24"
                          >
                            <path
                              stroke="currentColor"
                              strokeLinecap="round"
                              strokeLinejoin="round"
                              strokeWidth="1.5"
                              d="M7.757 16.243l8.486-8.486M16.243 16.243L7.757 7.757"
                            ></path>
                          </svg>
                        </Button>
                      )}
                    </div>
                  ))}
                <div
                  onClick={addNewTag}
                  className="flex cursor-pointer items-center justify-center gap-2 rounded-lg bg-white py-3 text-sm text-black-900/50 transition-colors hover:bg-komiser-700/10"
                >
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    width="16"
                    height="16"
                    fill="none"
                    viewBox="0 0 24 24"
                  >
                    <path
                      fill="currentColor"
                      d="M18 12.75H6c-.41 0-.75-.34-.75-.75s.34-.75.75-.75h12c.41 0 .75.34.75.75s-.34.75-.75.75z"
                    ></path>
                    <path
                      fill="currentColor"
                      d="M12 18.75c-.41 0-.75-.34-.75-.75V6c0-.41.34-.75.75-.75s.75.34.75.75v12c0 .41-.34.75-.75.75z"
                    ></path>
                  </svg>
                  Add new tag
                </div>
                <div className="flex items-center justify-end gap-6">
                  {((data && data.tags && data.tags.length > 0) ||
                    (!data && bulkItems)) && (
                    <Button
                      size="lg"
                      style="delete-ghost"
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
            )}

            {page === 'delete' && (
              <>
                <div className="flex flex-col gap-6">
                  <div className="mx-auto flex h-12 w-12 items-center justify-center rounded-full bg-error-100 text-error-600">
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
                    <p className="text-center font-medium text-black-900">
                      Are you sure you want to delete all tags from{' '}
                      {formatNumber(bulkItems.length)}{' '}
                      {bulkItems.length > 1 ? 'resources' : 'resource'}?
                    </p>
                    <p className="text-sm text-black-400">
                      This is a permanent action, and it will also delete
                      previous tags you have added to these resources.
                    </p>
                  </div>
                  <div className="flex items-center justify-end gap-6">
                    <Button
                      size="lg"
                      style="secondary"
                      onClick={() => {
                        goTo('tags');
                      }}
                    >
                      Cancel
                    </Button>
                    <Button
                      size="lg"
                      style="delete-ghost"
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
        </div>
      </Sidepanel>
    </>
  );
}

export default InventorySidePanel;
