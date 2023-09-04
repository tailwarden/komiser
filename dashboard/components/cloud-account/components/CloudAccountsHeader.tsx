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
          <p className="flex items-center gap-2 text-lg font-medium text-black-900">
            Your Cloud Accounts
          </p>
          <Button type="button" style="secondary" onClick={() => {}}>
            <PlusIcon className="-ml-2 mr-1 h-6 w-6" />
            Add Cloud Accounts
          </Button>
        </>
      )}
    </div>
  );
}

export default CloudAccountsHeader;
