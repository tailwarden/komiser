import { useState, useEffect, useRef } from 'react';
import Avatar from '@components/avatar/Avatar';
import platform from '@utils/providerHelper';
import { RefreshIcon } from '@components/icons';
import settingsService from '@services/settingsService';
import { CloudAccount } from '../hooks/useCloudAccounts/useCloudAccount';
import CloudAccountStatus from './CloudAccountStatus';
import More2Icon from '../../icons/More2Icon';
import DeleteIcon from '../../icons/DeleteIcon';
import EditIcon from '../../icons/EditIcon';

export default function CloudAccountItem({
  account,
  openModal,
  setEditCloudAccount,
  setIsDeleteModalOpen,
  setCloudAccountItem
}: {
  account: CloudAccount;
  openModal: (cloudAccount: CloudAccount) => void;
  setEditCloudAccount: (editCloudAccount: boolean) => void;
  setIsDeleteModalOpen: (isDeleteModalOpen: boolean) => void;
  setCloudAccountItem: (cloudAccountItem: CloudAccount) => void;
}) {
  const optionsRef = useRef<HTMLDivElement | null>(null);
  const { id, provider: cloudProvider, name, status } = account;
  const [isOpen, setIsOpen] = useState(false);

  useEffect(() => {
    const handleOutsideClick = (event: MouseEvent) => {
      if (
        optionsRef.current &&
        !optionsRef.current.contains(event.target as Node)
      ) {
        setIsOpen(false); // Close the options if clicked outside
      }
    };

    document.addEventListener('mousedown', handleOutsideClick);

    return () => {
      document.removeEventListener('mousedown', handleOutsideClick);
    };
  }, []);

  const handleRescanClick = () => {
    settingsService.rescanCloudAccount(id as number).then(res => {
      if (res === Error) {
        console.log('error', res);
      } else {
        window.location.reload();
      }
    });
  };

  return (
    <div
      key={id}
      onClick={() => openModal(account)}
      className="relative my-5 flex w-full items-center gap-4 rounded-lg border-2 border-gray-200 bg-white p-6 text-gray-950 transition-colors"
    >
      <Avatar avatarName={cloudProvider} size={48} />
      <div className="mr-auto">
        <p className="font-bold">{name}</p>
        <p className="text-gray-500">{platform.getLabel(cloudProvider)}</p>
      </div>

      <CloudAccountStatus status={status} />

      <More2Icon
        className="h-6 w-6 cursor-pointer"
        onClick={() => setIsOpen(!isOpen)}
      />

      {isOpen && (
        <div
          ref={optionsRef}
          className="absolute right-0 top-0 mr-5 mt-[70px] items-center rounded-md border border-gray-100 bg-white p-4 shadow-right"
          style={{ zIndex: 1000 }}
        >
          <button
            className="flex w-full rounded-md py-3 pl-3 pr-5 text-left text-sm text-gray-700 hover:bg-background-ds"
            onClick={() => {
              setIsOpen(false);
              setEditCloudAccount(true);
            }}
          >
            <EditIcon className="mr-2 h-6 w-6" />
            Edit cloud account
          </button>
          <button
            className="flex w-full rounded-md py-3 pl-3 pr-5 text-left text-sm text-gray-700 hover:bg-background-ds"
            onClick={() => {
              handleRescanClick();
            }}
            disabled={status === 'SCANNING'}
            style={{
              opacity: status === 'SCANNING' ? 0.5 : 1,
              pointerEvents: status === 'SCANNING' ? 'none' : 'auto'
            }}
          >
            <RefreshIcon className="mr-2 h-6 w-6" />
            Rescan
          </button>
          <button
            className="flex w-full rounded-md py-3 pl-3 pr-5 text-left text-sm text-red-500 hover:bg-background-ds"
            onClick={() => {
              setIsOpen(false);
              setIsDeleteModalOpen(true);
              setCloudAccountItem(account);
            }}
          >
            <DeleteIcon className="mr-2 h-6 w-6" />
            Remove account
          </button>
        </div>
      )}
    </div>
  );
}
