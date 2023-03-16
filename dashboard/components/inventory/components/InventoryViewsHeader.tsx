import type { NextRouter } from 'next/router';
import type { FormEvent } from 'react';
import { useState } from 'react';
import Button from '@/components/button/Button';
import Modal from '@/components/modal/Modal';
import type { ToastProps } from '@/components/toast/hooks/useToast';
import type {
  InventoryFilterData,
  View
} from '../hooks/useInventory/types/useInventoryTypes';
import type { ViewsPages } from './view/hooks/useViews';

type InventoryViewsHeaderProps = {
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
  setToast: (toast: ToastProps | undefined) => void;
};

function InventoryViewsHeader({
  openModal,
  views,
  router,
  saveView,
  loading,
  deleteView,
  deleteLoading,
  setToast
}: InventoryViewsHeaderProps) {
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
    <div className="relative">
      {currentView && (
        <div className="flex items-center gap-2 text-lg font-medium text-black-900">
          <span>{currentView.name}</span>
          <Button style="ghost" size="xs" onClick={openDropdown}>
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
                strokeMiterlimit="10"
                strokeWidth="1.5"
                d="M19.92 8.95l-6.52 6.52c-.77.77-2.03.77-2.8 0L4.08 8.95"
              ></path>
            </svg>
          </Button>
        </div>
      )}

      {dropdownIsOpen && (
        <>
          <div
            onClick={closeDropdown}
            className="fixed inset-0 z-20 hidden animate-fade-in bg-transparent opacity-0 sm:block"
          ></div>
          <div className="absolute left-0 top-10 z-[21] inline-flex w-[16rem] rounded-lg bg-white p-4 text-sm shadow-xl">
            <div className="flex w-full flex-col gap-1">
              <Button
                style="ghost"
                size="sm"
                align="left"
                gap="md"
                transition={false}
                onClick={() => {
                  closeDropdown();
                  openModal();
                }}
              >
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
                    d="M11 2H9C4 2 2 4 2 9v6c0 5 2 7 7 7h6c5 0 7-2 7-7v-2"
                  ></path>
                  <path
                    stroke="currentColor"
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeMiterlimit="10"
                    strokeWidth="1.5"
                    d="M16.04 3.02L8.16 10.9c-.3.3-.6.89-.66 1.32l-.43 3.01c-.16 1.09.61 1.85 1.7 1.7l3.01-.43c.42-.06 1.01-.36 1.32-.66l7.88-7.88c1.36-1.36 2-2.94 0-4.94-2-2-3.58-1.36-4.94 0zM14.91 4.15a7.144 7.144 0 004.94 4.94"
                  ></path>
                </svg>
                Edit view
              </Button>
              <Button
                style="ghost"
                size="sm"
                align="left"
                gap="md"
                transition={false}
                onClick={e => {
                  closeDropdown();
                  saveView(e, true, currentView);
                }}
                loading={loading}
              >
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
                    d="M16 12.9v4.2c0 3.5-1.4 4.9-4.9 4.9H6.9C3.4 22 2 20.6 2 17.1v-4.2C2 9.4 3.4 8 6.9 8h4.2c3.5 0 4.9 1.4 4.9 4.9z"
                  ></path>
                  <path
                    stroke="currentColor"
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth="1.5"
                    d="M22 6.9v4.2c0 3.5-1.4 4.9-4.9 4.9H16v-3.1C16 9.4 14.6 8 11.1 8H8V6.9C8 3.4 9.4 2 12.9 2h4.2C20.6 2 22 3.4 22 6.9z"
                  ></path>
                </svg>
                Duplicate view
              </Button>
              <Button
                style="ghost"
                size="sm"
                align="left"
                gap="md"
                transition={false}
                onClick={() => {
                  closeDropdown();
                  openModal(undefined, 'alerts');
                }}
                loading={loading}
              >
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
                    strokeMiterlimit="10"
                    strokeWidth="1.5"
                    d="M12.02 2.91c-3.31 0-6 2.69-6 6v2.89c0 .61-.26 1.54-.57 2.06L4.3 15.77c-.71 1.18-.22 2.49 1.08 2.93 4.31 1.44 8.96 1.44 13.27 0 1.21-.4 1.74-1.83 1.08-2.93l-1.15-1.91c-.3-.52-.56-1.45-.56-2.06V8.91c0-3.3-2.7-6-6-6z"
                  ></path>
                  <path
                    stroke="currentColor"
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeMiterlimit="10"
                    strokeWidth="1.5"
                    d="M13.87 3.2a6.754 6.754 0 00-3.7 0c.29-.74 1.01-1.26 1.85-1.26.84 0 1.56.52 1.85 1.26z"
                  ></path>
                  <path
                    stroke="currentColor"
                    strokeMiterlimit="10"
                    strokeWidth="1.5"
                    d="M15.02 19.06c0 1.65-1.35 3-3 3-.82 0-1.58-.34-2.12-.88a3.01 3.01 0 01-.88-2.12"
                  ></path>
                </svg>
                Set up alert
              </Button>
              <Button
                style="ghost"
                size="sm"
                align="left"
                gap="md"
                transition={false}
                onClick={() => {
                  navigator.clipboard.writeText(document.URL);
                  setToast({
                    hasError: false,
                    title: 'Link copied!',
                    message: `${document.URL} has been copied to your clipboard.`
                  });
                }}
              >
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
                    d="M13.06 10.94a5.74 5.74 0 010 8.13c-2.25 2.24-5.89 2.25-8.13 0-2.24-2.25-2.25-5.89 0-8.13"
                  ></path>
                  <path
                    stroke="currentColor"
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth="1.5"
                    d="M10.59 13.41c-2.34-2.34-2.34-6.14 0-8.49 2.34-2.35 6.14-2.34 8.49 0 2.35 2.34 2.34 6.14 0 8.49"
                  ></path>
                </svg>
                Copy view link
              </Button>
              <span className="m-2 -mx-4 border-b border-black-200/40"></span>
              <Button
                style="ghost"
                size="sm"
                align="left"
                gap="md"
                transition={false}
                onClick={() => {
                  closeDropdown();
                  openDoubleConfirmationModal();
                }}
              >
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
                    d="M21 5.98c-3.33-.33-6.68-.5-10.02-.5-1.98 0-3.96.1-5.94.3L3 5.98M8.5 4.97l.22-1.31C8.88 2.71 9 2 10.69 2h2.62c1.69 0 1.82.75 1.97 1.67l.22 1.3M18.85 9.14l-.65 10.07C18.09 20.78 18 22 15.21 22H8.79C6 22 5.91 20.78 5.8 19.21L5.15 9.14M10.33 16.5h3.33M9.5 12.5h5"
                  ></path>
                </svg>
                Delete view
              </Button>
            </div>
          </div>
        </>
      )}

      <Modal isOpen={modalIsOpen} closeModal={closeDoubleConfirmationModal}>
        <div className="flex w-full flex-col gap-6 rounded-lg">
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
                  strokeWidth="2"
                  d="M12 22c5.5 0 10-4.5 10-10S17.5 2 12 2 2 6.5 2 12s4.5 10 10 10zM12 8v5M11.995 16h.009"
                ></path>
              </svg>
            </div>
            <div className="flex flex-col items-center gap-6">
              <p className="text-center font-medium text-black-900">
                Are you sure you want to delete this view?
              </p>
              <p className="text-sm text-black-400">
                This is a permanent action.
              </p>
            </div>
          </div>
          <div className="flex flex-wrap items-center justify-end gap-6">
            <Button
              type="button"
              size="lg"
              style="ghost"
              onClick={closeDoubleConfirmationModal}
            >
              Cancel
            </Button>
            <Button
              type="button"
              size="lg"
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
    </div>
  );
}

export default InventoryViewsHeader;
