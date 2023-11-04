import { useRouter } from 'next/router';

import EmptyState from '@components/empty-state/EmptyState';
import HyperLinkIcon from '@components/icons/HyperLinkIcon';
import useDependencyGraph from '../hooks/useDependencyGraph';
import SingleDependencyGraphLoader from './SingleDependencyGraphLoader';

function SingleDependencyGraphWrapper({ resourceId }: { resourceId?: string }) {
  const { loading, data, error, fetch } = useDependencyGraph(resourceId);
  const router = useRouter();
  const title = 'Open in explorer';
  return (
    <>
      <div className="flex h-[calc(100vh-145px)] w-full flex-col">
        <div className="flex flex-row justify-between gap-2">
          <div className=" flex w-48 items-center gap-1 truncate text-base font-medium leading-normal text-primary">
            <a
              target="_blank"
              onClick={() => router.push('/explorer')}
              rel="noreferrer"
              className="hover:text-primary"
            >
              <HyperLinkIcon />
            </a>
            <span>{title}</span>
          </div>
        </div>
        {!data?.nodes.length && !data?.edges.length ? (
          <div className="mt-24">
            <EmptyState
              title="We could not find any resources"
              message="It seems like you have no AWS cloud resources associated with your cloud accounts"
              mascotPose="devops"
              secondaryActionLabel="Report an issue"
              actionLabel="Check cloud account"
              secondaryAction={() => {
                router.push(
                  'https://github.com/tailwarden/komiser/issues/new/choose'
                );
              }}
              action={() => {
                router.push('/');
              }}
            />
          </div>
        ) : (
          <SingleDependencyGraphLoader
            loading={loading}
            data={data}
            error={error}
            fetch={fetch}
          />
        )}
      </div>
    </>
  );
}

export default SingleDependencyGraphWrapper;
