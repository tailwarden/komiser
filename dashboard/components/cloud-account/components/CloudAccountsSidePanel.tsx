import { FormEvent, useState } from 'react';
import AzureAccountDetails from '@components/account-details/AzureAccountDetails';
import GcpAccountDetails from '@components/account-details/GcpAccountDetails';
import DigitalOceanAccountDetails from '@components/account-details/DigitalOceanAccountDetails';
import CivoAccountDetails from '@components/account-details/CivoAccountDetails';
import LinodeAccountDetails from '@components/account-details/LinodeAccountDetails';
import KubernetesAccountDetails from '@components/account-details/KubernetesAccountDetails';
import TencentAccountDetails from '@components/account-details/TencentAccountDetails';
import MongoDbAtlasAccountDetails from '@components/account-details/MongoDBAtlasAccountDetails';
import OciAccountDetails from '@components/account-details/OciAccountDetails';
import ScalewayAccountDetails from '@components/account-details/ScalewayAccountDetails';
import { getPayloadFromForm } from '@utils/cloudAccountHelpers';
import { ToastProps } from '@components/toast/Toast';
import Avatar from '@components/avatar/Avatar';
import { allProviders, Provider } from '../../../utils/providerHelper';
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
import settingsService from '../../../services/settingsService';

interface CloudAccountsSidePanelProps {
  isOpen: boolean;
  closeModal: () => void;
  cloudAccount: CloudAccount;
  cloudAccounts: CloudAccount[];
  setCloudAccounts: (cloudAccounts: CloudAccount[]) => void;
  handleAfterDelete: (account: CloudAccount) => void;
  page: CloudAccountsPage;
  goTo: (page: CloudAccountsPage) => void;
  showToast: (toast: ToastProps) => void;
}

function AccountDetails({
  cloudAccountData
}: {
  cloudAccountData: CloudAccount;
}) {
  switch (cloudAccountData.provider.toLocaleLowerCase()) {
    case allProviders.AWS:
      return <AwsAccountDetails cloudAccountData={cloudAccountData} />;
    case allProviders.GCP:
      return <GcpAccountDetails cloudAccountData={cloudAccountData} />;
    case allProviders.DIGITAL_OCEAN:
      return <DigitalOceanAccountDetails cloudAccountData={cloudAccountData} />;
    case allProviders.AZURE:
      return <AzureAccountDetails cloudAccountData={cloudAccountData} />;
    case allProviders.CIVO:
      return <CivoAccountDetails cloudAccountData={cloudAccountData} />;
    case allProviders.KUBERNETES:
      return <KubernetesAccountDetails cloudAccountData={cloudAccountData} />;
    case allProviders.LINODE:
      return <LinodeAccountDetails cloudAccountData={cloudAccountData} />;
    case allProviders.TENCENT:
      return <TencentAccountDetails cloudAccountData={cloudAccountData} />;
    case allProviders.OCI:
      return <OciAccountDetails cloudAccountData={cloudAccountData} />;
    case allProviders.SCALE_WAY:
      return <ScalewayAccountDetails cloudAccountData={cloudAccountData} />;
    case allProviders.MONGODB_ATLAS:
      return <MongoDbAtlasAccountDetails cloudAccountData={cloudAccountData} />;
    default:
      return null;
  }
}

function CloudAccountsSidePanel({
  isOpen,
  closeModal,
  cloudAccount,
  cloudAccounts,
  setCloudAccounts,
  handleAfterDelete,
  page,
  goTo,
  showToast
}: CloudAccountsSidePanelProps) {
  const [isDeleteOpen, setIsDeleteOpen] = useState(false);
  const [loading, setLoading] = useState(false);

  const handleEditCloudAccount = (
    event: FormEvent<HTMLFormElement>,
    id: number | undefined,
    provider: Provider
  ) => {
    event.preventDefault();
    if (!id) return false;

    setLoading(true);
    const payloadJson = JSON.stringify(
      getPayloadFromForm(new FormData(event.currentTarget), provider)
    );
    settingsService.editCloudAccount(id, payloadJson).then(res => {
      if (res === Error || res.error) {
        setLoading(false);
        showToast({
          hasError: true,
          title: 'Cloud account not edited',
          message:
            'There was an error editing this cloud account. Refer to the logs and try again.'
        });
      } else {
        setLoading(false);
        showToast({
          hasError: false,
          title: 'Cloud account edited',
          message: `The cloud account was successfully edited!`
        });
        setCloudAccounts(
          cloudAccounts.map(c => (c.id === cloudAccount.id ? res : c))
        );
        closeModal();
      }
    });

    return true;
  };

  return (
    <>
      <Sidepanel isOpen={isOpen} closeModal={closeModal}>
        <div className="flex h-full flex-col">
          {/* Modal headers */}
          <div className="flex flex-wrap-reverse items-center justify-between gap-6 sm:flex-nowrap">
            {cloudAccount && (
              <div className="flex flex-wrap items-center gap-4 sm:flex-nowrap">
                <Avatar avatarName={cloudAccount.provider} size={40} />
                <div className="flex flex-col gap-1">
                  <div className="flex max-w-[14rem] items-center gap-1">
                    <p className="truncate font-medium text-gray-950">
                      {cloudAccount.name}
                    </p>
                  </div>
                  <p className="flex items-center gap-2 text-xs text-gray-500">
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
            <div className="mt-4 flex flex-col gap-6 rounded-lg bg-gray-50 p-8">
              <CloudAccountDeleteContents
                cloudAccount={cloudAccount}
                onCancel={() => setIsDeleteOpen(false)}
                handleAfterDelete={(account: CloudAccount) => {
                  handleAfterDelete(account);
                  setIsDeleteOpen(false);
                  closeModal();
                }}
                showToast={showToast}
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
                <form
                  onSubmit={event =>
                    handleEditCloudAccount(
                      event,
                      cloudAccount.id,
                      cloudAccount.provider
                    )
                  }
                >
                  <input type="hidden" name="test" value="test" />
                  <div className="-mx-4 min-h-0 flex-grow overflow-y-auto px-4">
                    <label className="mb-2 mt-6 block text-gray-700">
                      Status
                    </label>
                    <CloudAccountStatus status={cloudAccount?.status} />
                    <AccountDetails cloudAccountData={cloudAccount} />
                  </div>
                  <div className="-mx-4 mb-4 h-px bg-gray-300"></div>
                  <div className="mb-4 flex flex-shrink-0 items-center justify-end gap-6">
                    <Button size="lg" style="ghost" onClick={closeModal}>
                      Cancel
                    </Button>
                    <Button size="lg" type="submit" loading={loading}>
                      Save changes
                    </Button>
                  </div>
                </form>
              )}
            </>
          )}
        </div>
      </Sidepanel>
    </>
  );
}

export default CloudAccountsSidePanel;
