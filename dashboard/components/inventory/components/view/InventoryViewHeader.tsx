import { ToastProps } from '@components/toast/Toast';
import { NextRouter } from 'next/router';
import { FormEvent, useState } from 'react';
import Button from '../../../button/Button';
import ExportCSV from '../../../export-csv/ExportCSV';
import AlertIcon from '../../../icons/AlertIcon';
import ChevronDownIcon from '../../../icons/ChevronDownIcon';
import DeleteIcon from '../../../icons/DeleteIcon';
import DuplicateIcon from '../../../icons/DuplicateIcon';
import EditIcon from '../../../icons/EditIcon';
import LinkIcon from '../../../icons/LinkIcon';
import WarningIcon from '../../../icons/WarningIcon';
import Modal from '../../../modal/Modal';
import {
  InventoryFilterData,
  View
} from '../../hooks/useInventory/types/useInventoryTypes';
import { ViewsPages } from './hooks/useViews';

type InventoryViewHeaderProps = {
  openModal: (
    filters?: InventoryFilterData[],
    openPage?: ViewsPages | undefined
  ) => void;
  views: View[] | undefined;
  router: NextRouter;
  saveView: (
    e: FormEvent<HTMLFormElement>,
    duplicate?: boolean | undefined,
    viewToBeDuplicated?: View | undefined
  ) => void;
  loading: boolean;
  deleteLoading: boolean;
  deleteView: (
    dropdown?: boolean | undefined,
    viewToBeDeleted?: View | undefined
  ) => void;
  showToast: (toast: ToastProps) => void;
};

function InventoryViewHeader({
  openModal,
  views,
  router,
  saveView,
  loading,
  deleteView,
  deleteLoading,
  showToast
}: InventoryViewHeaderProps) {
  const [dropdownIsOpen, setDropdownIsOpen] = useState(false);
  const [modalIsOpen, setModalIsOpen] = useState(false);

  function closeDropdown() {
    setDropdownIsOpen(false);
  }

  function openDropdown() {
    setDropdownIsOpen(true);
  }

  function closeDoubleConfirmationModal() {
    setModalIsOpen(false);
  }

  function openDoubleConfirmationModal() {
    setModalIsOpen(true);
  }

  const currentView = views?.find(
    view => view.id.toString() === router.query.view
  );

  return (
    <div className="absolute -top-14 left-0">
      {currentView && (
        <>
          <div className="flex items-center gap-2 text-lg font-medium text-gray-950">
            <span>{currentView.name}</span>
            <Button style="ghost" size="xxs" onClick={openDropdown}>
              <ChevronDownIcon width={16} height={16} />
            </Button>
          </div>

          {dropdownIsOpen && (
            <>
              <div
                onClick={closeDropdown}
                className="fixed inset-0 z-20 hidden animate-fade-in bg-transparent opacity-0 sm:block"
              ></div>
              <div className="absolute left-0 top-10 z-[21] inline-flex w-[16rem] rounded-lg bg-white p-4 text-sm shadow-right">
                <div className="flex w-full flex-col gap-1">
                  <Button
                    style="dropdown"
                    size="sm"
                    transition={false}
                    onClick={() => {
                      closeDropdown();
                      openModal();
                    }}
                  >
                    <EditIcon width={24} height={24} />
                    Edit view
                  </Button>
                  <Button
                    style="dropdown"
                    size="sm"
                    transition={false}
                    onClick={e => {
                      closeDropdown();
                      saveView(e, true, currentView);
                    }}
                    loading={loading}
                  >
                    <DuplicateIcon width={24} height={24} />
                    Duplicate view
                  </Button>
                  <Button
                    style="dropdown"
                    size="sm"
                    transition={false}
                    onClick={() => {
                      closeDropdown();
                      openModal(undefined, 'alerts');
                    }}
                    loading={loading}
                  >
                    <AlertIcon width={24} height={24} />
                    Set up alert
                  </Button>
                  <Button
                    style="dropdown"
                    size="sm"
                    transition={false}
                    onClick={() => {
                      navigator.clipboard.writeText(document.URL);
                      showToast({
                        hasError: false,
                        title: 'Link copied!',
                        message: `${document.URL} has been copied to your clipboard.`
                      });
                    }}
                  >
                    <LinkIcon width={24} height={24} />
                    Copy view link
                  </Button>
                  <ExportCSV showToast={showToast} />
                  <span className="m-2 -mx-4 border-b border-gray-300"></span>
                  <Button
                    style="dropdown"
                    size="sm"
                    transition={false}
                    onClick={() => {
                      closeDropdown();
                      openDoubleConfirmationModal();
                    }}
                  >
                    <DeleteIcon width={24} height={24} />
                    Delete view
                  </Button>
                </div>
              </div>
            </>
          )}

          <Modal isOpen={modalIsOpen} closeModal={closeDoubleConfirmationModal}>
            <div className="flex w-full flex-col gap-6 rounded-lg">
              <div className="flex flex-col gap-6">
                <div className="mx-auto flex h-12 w-12 items-center justify-center rounded-full bg-red-50 text-red-500">
                  <WarningIcon width={24} height={24} />
                </div>
                <div className="flex flex-col items-center gap-2">
                  <p className="text-center text-lg font-medium text-gray-950">
                    Are you sure you want to delete this view?
                  </p>
                  <p className="text-sm text-gray-700">
                    This is a permanent action.
                  </p>
                </div>
              </div>
              <div className="flex flex-wrap items-center justify-end gap-6">
                <Button
                  type="button"
                  style="ghost"
                  onClick={closeDoubleConfirmationModal}
                >
                  Cancel
                </Button>
                <Button
                  type="button"
                  style="delete"
                  onClick={() => {
                    deleteView(true, currentView);
                  }}
                  loading={deleteLoading}
                >
                  {deleteLoading ? 'Deleting...' : 'Delete view'}
                </Button>
              </div>
            </div>
          </Modal>
        </>
      )}
    </div>
  );
}

export default InventoryViewHeader;
