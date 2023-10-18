import { useState } from 'react';
import router from 'next/router';
import Head from 'next/head';
import Link from 'next/link';
import Image from 'next/image';
import OnboardingWizardLayout, {
  LeftSideLayout,
  RightSideLayout
} from '@components/onboarding-wizard/OnboardingWizardLayout';
import PlusIcon from '@components/icons/PlusIcon';
import providers from '@utils/providerHelper';
import DeleteIcon from '@components/icons/DeleteIcon';
import Modal from '@components/modal/Modal';
import CloudAccountDeleteContents from '@components/cloud-account/components/CloudAccountDeleteContents';
import Toast from '@components/toast/Toast';

import useCloudAccount from '@components/cloud-account/hooks/useCloudAccounts/useCloudAccount';
import Button from '@components/button/Button';
import { useToast } from '@components/toast/ToastProvider';

export default function CloudAccounts() {
  const [isDeleteModalOpen, setIsDeleteModalOpen] = useState<boolean>(false);

  const { toast, showToast, dismissToast } = useToast();

  const {
    cloudAccounts,
    setCloudAccounts,
    cloudAccountItem,
    setCloudAccountItem
  } = useCloudAccount();

  const closeRemoveModal = () => {
    setIsDeleteModalOpen(false);
  };

  const handleDelete = (account: any) => {
    setCloudAccountItem(account);
    setIsDeleteModalOpen(true);
  };

  const handleAfterDelete = (account: any) => {
    setCloudAccounts(
      cloudAccounts.filter((item: any) => item.id !== account.id)
    );
    setIsDeleteModalOpen(false);
  };

  return (
    <div>
      <Head>
        <title>Onboarding - Komiser</title>
        <meta name="description" content="Onboarding - Komiser" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <OnboardingWizardLayout>
        <LeftSideLayout title="Connected cloud accounts" progressBarWidth="35%">
          <div className="mb-8 leading-6 text-gray-900/60">
            <div className="font-normal">
              Here, you can add more cloud accounts or edit/delete existing
              ones, before moving to the next step.
            </div>
          </div>
          <div className="mb-4 space-y-4">
            <Link
              href={'/onboarding/choose-cloud/'}
              className="flex w-full items-center rounded border-[1.5px] border-primary bg-transparent p-6 text-primary hover:bg-komiser-130"
            >
              <PlusIcon className="my-4 ml-2 mr-6 h-6 w-6" />
              Add cloud account
            </Link>
            {cloudAccounts.map(account => (
              <div
                key={account.id}
                className="flex items-center justify-between rounded-lg border border-black-200 p-6"
              >
                <div className="flex flex-wrap items-center gap-4 sm:flex-nowrap">
                  <picture className="flex-shrink-0">
                    <Image
                      src={String(providers.providerImg(account.provider))}
                      className="rounded-full"
                      height={40}
                      width={40}
                      alt={account.provider}
                    />
                  </picture>

                  <div className="flex flex-col gap-1">
                    <div className="flex max-w-[14rem] items-center gap-1">
                      <p className="truncate font-medium text-black-900">
                        {account.name}
                      </p>
                    </div>
                    <p className="flex items-center gap-2 text-xs text-black-300">
                      {providers.providerLabel(account.provider)}
                    </p>
                  </div>
                </div>
                <div className="flex gap-5">
                  <button
                    className="hidden items-center gap-2 transition-colors hover:text-primary md:flex"
                    onClick={() => handleDelete(account)}
                  >
                    <DeleteIcon className="h-4 w-4" />
                  </button>
                </div>
              </div>
            ))}
          </div>
          <div className="fixed bottom-0 -mx-20 flex w-[calc(100%*6/11)] justify-end border-t border-black-200 bg-white px-20 py-4">
            <Button onClick={() => router.push('/onboarding/choose-database')}>
              Next
            </Button>
          </div>
        </LeftSideLayout>

        <RightSideLayout isCustom={true} customClasses="flex flex-col p-4">
          <div className="grid w-full grid-cols-7 gap-3">
            {/* Row 1 */}
            <div className="aspect-square h-full w-full rounded-lg bg-transparent"></div>
            <div className="aspect-square h-full w-full rounded-lg bg-transparent"></div>
            <div className="aspect-square h-full w-full rounded-lg bg-transparent"></div>
            <div className="aspect-square h-full w-full rounded-lg bg-transparent"></div>
            <div className="aspect-square h-full w-full rounded-lg bg-komiser-300"></div>
            <div className="aspect-square h-full w-full rounded-lg bg-transparent"></div>
            <div className="aspect-square h-full w-full rounded-lg bg-transparent"></div>

            {/* Row 2 */}
            <div className="aspect-square h-full w-full rounded-lg bg-transparent"></div>
            <div className="aspect-square h-full w-full rounded-lg bg-komiser-200"></div>
            <div className="aspect-square h-full w-full rounded-lg bg-komiser-200">
              <div className="relative bottom-3 left-3 h-full w-full scale-110 overflow-clip rounded-lg shadow-xl">
                <Image
                  src={String(providers.providerImg('aws'))}
                  layout="fill"
                  objectFit="cover"
                  className="object-center"
                  alt="AWS"
                />
              </div>
            </div>
            <div className="aspect-square h-full w-full rounded-lg bg-komiser-200"></div>
            <div className="aspect-square h-full w-full rounded-lg bg-komiser-200"></div>
            <div className="aspect-square h-full w-full rounded-lg bg-transparent"></div>
            <div className="aspect-square h-full w-full rounded-lg bg-transparent"></div>

            {/* Row 3 */}
            <div className="aspect-square h-full w-full rounded-lg bg-komiser-200"></div>
            <div className="aspect-square h-full w-full rounded-lg bg-komiser-200">
              <div className="relative h-full w-full overflow-clip rounded-lg shadow-xl">
                <Image
                  src={String(providers.providerImg('civo'))}
                  layout="fill"
                  objectFit="cover"
                  className="object-center"
                  alt="Civo"
                />
              </div>
            </div>
            <div className="aspect-square h-full w-full rounded-lg bg-komiser-200"></div>
            <div className="aspect-square h-full w-full rounded-lg bg-komiser-200"></div>
            <div className="aspect-square h-full w-full rounded-lg bg-transparent"></div>
            <div className="aspect-square h-full w-full rounded-lg bg-transparent"></div>
            <div className="aspect-square h-full w-full rounded-lg bg-komiser-400"></div>

            {/* Row 4 */}
            <div className="aspect-square h-full w-full rounded-lg bg-transparent"></div>
            <div className="aspect-square h-full w-full rounded-lg bg-komiser-200"></div>
            <div className="aspect-square h-full w-full rounded-lg bg-komiser-400">
              <div className="relative h-full w-full overflow-clip rounded-lg bg-white shadow-xl">
                <Image
                  src={String(providers.providerImg('gcp'))}
                  layout="fill"
                  objectFit="cover"
                  className="object-center p-2"
                  alt="GCP"
                />
              </div>
            </div>
            <div className="aspect-square h-full w-full rounded-lg bg-transparent"></div>
            <div className="aspect-square h-full w-full rounded-lg bg-transparent"></div>
            <div className="aspect-square h-full w-full rounded-lg bg-komiser-200"></div>
            <div className="aspect-square h-full w-full rounded-lg bg-transparent"></div>

            {/* Row 5 */}
            <div className="aspect-square h-full w-full rounded-lg bg-komiser-200"></div>
            <div className="aspect-square h-full w-full rounded-lg bg-komiser-200"></div>
            <div className="aspect-square h-full w-full rounded-lg bg-komiser-400"></div>
            <div className="aspect-square h-full w-full rounded-lg bg-komiser-200">
              <div className="relative left-3 top-3  h-full w-full overflow-clip rounded-lg bg-black-800 shadow-xl">
                <Image
                  src={String(providers.providerImg('azure'))}
                  layout="fill"
                  objectFit="cover"
                  className="object-center p-5"
                  alt="Azure"
                />
              </div>
            </div>
            <div className="aspect-square h-full w-full rounded-lg bg-transparent"></div>
            <div className="aspect-square h-full w-full rounded-lg bg-komiser-200"></div>
            <div className="aspect-square h-full w-full rounded-lg bg-transparent"></div>

            {/* Row 6 */}
            <div className="aspect-square h-full w-full rounded-lg bg-transparent"></div>
            <div className="aspect-square h-full w-full rounded-lg bg-komiser-400"></div>
            <div className="aspect-square h-full w-full rounded-lg bg-komiser-200"></div>
            <div className="aspect-square h-full w-full rounded-lg bg-transparent"></div>
            <div className="aspect-square h-full w-full rounded-lg bg-komiser-200"></div>
            <div className="aspect-square h-full w-full rounded-lg bg-transparent"></div>
            <div className="aspect-square h-full w-full rounded-lg bg-transparent"></div>

            {/* Row 7 */}
            <div className="aspect-square h-full w-full rounded-lg bg-transparent"></div>
            <div className="aspect-square h-full w-full rounded-lg bg-transparent"></div>
            <div className="aspect-square h-full w-full rounded-lg bg-transparent"></div>
            <div className="aspect-square h-full w-full rounded-lg bg-komiser-200"></div>
            <div className="aspect-square h-full w-full rounded-lg bg-transparent"></div>
            <div className="aspect-square h-full w-full rounded-lg bg-transparent"></div>
            <div className="aspect-square h-full w-full rounded-lg bg-transparent"></div>
          </div>
        </RightSideLayout>
      </OnboardingWizardLayout>

      <Modal isOpen={isDeleteModalOpen} closeModal={() => closeRemoveModal()}>
        <div className="flex max-w-xl flex-col gap-y-6 p-8 text-black-400">
          {cloudAccountItem && (
            <CloudAccountDeleteContents
              cloudAccount={cloudAccountItem}
              onCancel={closeRemoveModal}
              showToast={showToast}
              handleAfterDelete={handleAfterDelete}
            />
          )}
        </div>
      </Modal>
    </div>
  );
}
