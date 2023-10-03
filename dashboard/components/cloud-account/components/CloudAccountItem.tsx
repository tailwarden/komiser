import { useState, useEffect, useRef } from 'react';
import Image from 'next/image';
import providers from '@utils/providerHelper';
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
  const { id, provider, name, status } = account;
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

  return (
    <div
      key={id}
      onClick={() => openModal(account)}
      className="relative my-5 flex w-full items-center gap-4 rounded-lg border-2 border-black-170 bg-white p-6 text-black-900 transition-colors"
    >
      <Image
        src={providers.providerImg(provider) as string}
        alt={`${name} image`}
        width={150}
        height={150}
        className="h-12 w-12 rounded-full"
      />

      <div className="mr-auto">
        <p className="font-bold">{name}</p>
        <p className="text-black-300">{providers.providerLabel(provider)}</p>
      </div>

      <CloudAccountStatus status={status} />

      <More2Icon
        className="h-6 w-6 cursor-pointer"
        onClick={() => setIsOpen(!isOpen)}
      />

      {isOpen && (
        <div
          ref={optionsRef}
          className="absolute right-0 top-0 mr-5 mt-[70px] items-center rounded-md border border-black-130 bg-white p-4 shadow-xl"
          style={{ zIndex: 1000 }}
        >
          <button
            className="flex w-full rounded-md py-3 pl-3 pr-5 text-left text-sm text-black-400 hover:bg-black-150"
            onClick={() => {
              setIsOpen(false);
              setEditCloudAccount(true);
            }}
          >
            <EditIcon className="mr-2 h-6 w-6" />
            Edit cloud account
          </button>
          <button
            className="flex w-full rounded-md py-3 pl-3 pr-5 text-left text-sm text-error-600 hover:bg-black-150"
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
