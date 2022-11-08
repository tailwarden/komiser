import providers from '../../../utils/providerHelper';
import Button from '../../button/Button';
import { InventoryItem, Pages, Tag } from '../hooks/useInventory';
import TagWrapper from './TagWrapper';

type InventorySidePanelProps = {
  closeModal: () => void;
  data: InventoryItem;
  goTo: (page: Pages) => void;
  page: Pages;
  updateTags: (action?: 'delete') => void;
  tags: Tag[] | [] | undefined;
  handleChange: (newData: Partial<Tag>, id?: number) => void;
  removeTag: (id: number) => void;
  addNewTag: () => void;
  loading: boolean;
  deleteLoading: boolean;
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
  deleteLoading
}: InventorySidePanelProps) {
  return (
    <>
      <div
        onClick={closeModal}
        className="hidden sm:block fixed inset-0 z-30 bg-black-900/10 opacity-0 animate-fade-in"
      ></div>
      <div className="fixed overflow-auto inset-0 z-30 sm:top-4 sm:bottom-4 sm:right-4 sm:left-auto w-full sm:w-[32rem] p-6 sm:rounded-lg shadow-2xl opacity-0 animate-fade-in-up sm:animate-fade-in-left bg-white">
        {/* Modal headers */}
        <div className="flex flex-wrap-reverse sm:flex-nowrap items-center justify-between gap-6">
          {data && (
            <div className="flex flex-wrap sm:flex-nowrap items-center gap-4">
              <picture className="flex-shrink-0">
                <img
                  src={providers.providerImg(data.provider)}
                  className="w-8 h-8 rounded-full"
                  alt={data.provider}
                />
              </picture>

              <div className="flex flex-col gap-1 py-1">
                <p className="font-medium text-black-900 w-48 truncate ...">
                  {data.service}
                </p>
                <p className="text-xs text-black-300">{data.name}</p>
              </div>
            </div>
          )}

          <div className="flex gap-4 flex-shrink-0">
            <Button style="secondary" onClick={closeModal}>
              Close
            </Button>
          </div>
        </div>

        {/* Tabs */}
        <div className="mt-4"></div>
        <div className="text-sm font-medium text-center border-b-2 border-black-150 text-black-300">
          <ul className="flex justify-between sm:justify-start -mb-[2px]">
            <li className="mr-2">
              <a
                onClick={() => goTo('tags')}
                className={`select-none inline-block py-4 px-2 sm:p-4 rounded-t-lg border-b-2 border-transparent hover:text-komiser-700 cursor-pointer 
                      ${
                        (page === 'tags' || page === 'delete') &&
                        `text-komiser-600 border-komiser-600 hover:text-komiser-600`
                      }`}
              >
                Tags
              </a>
            </li>
          </ul>
        </div>

        {/* Tags form */}
        <div className="mt-6"></div>
        <div className="p-6 bg-black-100 rounded-lg">
          <div className="flex flex-col gap-6">
            {page === 'tags' && (
              <form
                onSubmit={e => {
                  e.preventDefault();

                  updateTags();
                }}
                className="flex flex-col gap-6"
              >
                {tags &&
                  tags.map((tag, id) => (
                    <div key={id} className="flex gap-6">
                      <TagWrapper
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
                              strokeWidth="2"
                              d="M7.757 16.243l8.486-8.486M16.243 16.243L7.757 7.757"
                            ></path>
                          </svg>
                        </Button>
                      )}
                    </div>
                  ))}
                <div
                  onClick={addNewTag}
                  className="flex items-center justify-center gap-2 py-3 bg-white hover:bg-komiser-700/10 rounded-lg text-black-900/50 text-sm transition-colors cursor-pointer"
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
                  {data && data.tags && data.tags.length > 0 && (
                    <Button
                      size="lg"
                      style="delete"
                      loading={deleteLoading}
                      onClick={() => {
                        updateTags('delete');
                      }}
                    >
                      Delete all tags
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
          </div>
        </div>
      </div>
    </>
  );
}

export default InventorySidePanel;
