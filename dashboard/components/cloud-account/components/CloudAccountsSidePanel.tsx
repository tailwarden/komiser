import { useState } from 'react';
import providers from '../../../utils/providerHelper';
import AwsAccountDetails from '../../account-details/AwsAccountDetails';
import Button from '../../button/Button';
import Sidepanel from '../../sidepanel/Sidepanel';
import SidepanelTabs from '../../sidepanel/SidepanelTabs';
import CloudAccountStatus from './CloudAccountStatus';
import CloudAccountDeleteContents from './CloudAccountDeleteContents';
import {
  CloudAccount,
  CloudAccountsPage
} from '../hooks/useCloudAccounts/useCloudAccount';
import { ToastProps } from '../../toast/hooks/useToast';
import settingsService from '../../../services/settingsService';

interface CloudAccountsSidePanelProps {
  isOpen: boolean;
  closeModal: () => void;
  cloudAccount: CloudAccount;
  cloudAccounts: CloudAccount[];
  setCloudAccounts: (cloudAccounts: CloudAccount[]) => void;
  page: CloudAccountsPage;
  goTo: (page: CloudAccountsPage) => void;
  setToast: (toast: ToastProps) => void;
}

function CloudAccountsSidePanel({
  isOpen,
  closeModal,
  cloudAccount,
  cloudAccounts,
  setCloudAccounts,
  page,
  goTo,
  setToast
}: CloudAccountsSidePanelProps) {
  const [isDeleteOpen, setIsDeleteOpen] = useState(false);
  const [loading, setLoading] = useState(false);
  const [cloudAccountData, setCloudAccountData] =
    useState<CloudAccount>(cloudAccount);

  const handleEditCloudAccount = () => {
    if (!cloudAccountData.id) return false;

    setLoading(true);
    const payloadJson = JSON.stringify(cloudAccountData);
    settingsService
      .editCloudAccount(cloudAccountData.id, payloadJson)
      .then(res => {
        if (res === Error || res.error) {
          setLoading(false);
          setToast({
            hasError: true,
            title: 'Cloud account not edited',
            message:
              'There was an error editing this cloud account. Refer to the logs and try again.'
          });
        } else {
          setLoading(false);
          setToast({
            hasError: false,
            title: 'Cloud account edited',
            message: `The cloud account was successfully edited!`
          });
          setCloudAccounts(
            cloudAccounts.map(c =>
              c.id === cloudAccountData.id ? cloudAccountData : c
            )
          );
          closeModal();
        }
      });

    return true;
  };

  return (
    <>
      <Sidepanel isOpen={isOpen} closeModal={closeModal}>
        <div className="flex max-h-full flex-col">
          {/* Modal headers */}
          <div className="flex flex-wrap-reverse items-center justify-between gap-6 sm:flex-nowrap">
            {cloudAccount && (
              <div className="flex flex-wrap items-center gap-4 sm:flex-nowrap">
                <picture className="flex-shrink-0">
                  <img
                    src={providers.providerImg(cloudAccount.provider)}
                    className="h-10 w-10 rounded-full"
                    alt={cloudAccount.provider}
                  />
                </picture>

                <div className="flex flex-col gap-1">
                  <div className="flex max-w-[14rem] items-center gap-1">
                    <p className="truncate font-medium text-black-900">
                      {cloudAccount.name}
                    </p>
                    {/* <a
                    target="_blank"
                    href={data.link}
                    rel="noreferrer"
                    className="cursor-pointer hover:text-primary"
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
                  </a> */}
                  </div>
                  <p className="flex items-center gap-2 text-xs text-black-300">
                    {cloudAccount.resources} resources in this cloud account
                  </p>
                </div>
              </div>
            )}

            <div className="flex flex-shrink-0 items-center gap-2">
              {!isDeleteOpen && (
                <Button style="delete" onClick={() => setIsDeleteOpen(true)}>
                  Remove account
                </Button>
              )}

              <Button
                style="secondary"
                onClick={() => {
                  setIsDeleteOpen(false);
                  closeModal();
                }}
              >
                Close
              </Button>
            </div>
          </div>

          {isDeleteOpen ? (
            <div className="mt-4 flex flex-col gap-6 rounded-lg bg-black-100 p-8">
              <CloudAccountDeleteContents
                cloudAccount={cloudAccount}
                onCancel={() => setIsDeleteOpen(false)}
                setToast={setToast}
              />
            </div>
          ) : (
            <>
              {/* Tabs */}
              <SidepanelTabs
                goTo={goTo}
                page={page}
                tabs={['Cloud account details']}
              />

              {/* Cloud account details */}

              {page === 'cloud account details' && (
                <>
                  <div className="-mx-4 min-h-0 overflow-y-auto px-4">
                    <label className="mb-2 mt-6 block text-gray-700">
                      Status
                    </label>

                    <CloudAccountStatus status={cloudAccount?.status} />
                    <AwsAccountDetails
                      cloudAccountData={cloudAccountData}
                      setCloudAccountData={setCloudAccountData}
                    />
                  </div>
                  <div className="-mx-4 mb-4 h-px bg-black-200"></div>
                  <div className="flex flex-shrink-0 items-center justify-end gap-6">
                    <Button size="lg" style="ghost" onClick={closeModal}>
                      Cancel
                    </Button>
                    <Button
                      size="lg"
                      loading={loading}
                      onClick={handleEditCloudAccount}
                    >
                      Save changes
                    </Button>
                  </div>
                </>
              )}
            </>
          )}
        </div>
      </Sidepanel>
    </>
  );
}

export default CloudAccountsSidePanel;
