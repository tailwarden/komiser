import Button from '../../button/Button';
import PlusIcon from '../../icons/PlusIcon';

type CloudAccountsHeaderProps = {
  isNotCustomView: boolean;
};

function CloudAccountsHeader({ isNotCustomView }: CloudAccountsHeaderProps) {
  return (
    <div className="flex min-h-[40px] items-center justify-between gap-8">
      {isNotCustomView && (
        <>
          <p className="flex items-center gap-2 text-lg font-medium text-gray-950">
            Your Cloud Accounts
          </p>
        </>
      )}
    </div>
  );
}

export default CloudAccountsHeader;
