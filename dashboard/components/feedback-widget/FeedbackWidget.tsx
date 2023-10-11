import { useState, useRef, useCallback, memo, SyntheticEvent } from 'react';
import { FileUploader } from 'react-drag-drop-files';
// eslint-disable-next-line import/no-extraneous-dependencies
import { toBlob } from 'html-to-image';
import Image from 'next/image';

import Modal from '@components/modal/Modal';
import Input from '@components/input/Input';
import settingsService from '@services/settingsService';
import Button from '@components/button/Button';
import useToast from '@components/toast/hooks/useToast';
import Toast from '@components/toast/Toast';

// We define the placeholder here for convenience
// It's difficult to read when passed inline
const textAreaPlaceholder = `Example:
Steps to Reproduce
1. Describe the actions you took leading to the bug.
2. Include any specific settings or options you selected.

Expected Behavior
1. Explain what you expected to happen.
2. Detail how the feature or function should work.

Outcome
1. Describe what actually happened.
2. Include any error messages or unexpected behavior.`;

const useFeedbackWidget = (defaultState: boolean = false) => {
  const [showFeedbackModel, setShowFeedbackModal] = useState(defaultState);

  const FILE_TYPES = ['JPG', 'PNG', 'GIF', 'TXT', 'LOG', 'MP4', 'AVI', 'MOV'];
  const FEEDBACK_MODAL_ID = 'feedback-modal';
  const MAX_FILE_SIZE_MB = 37;

  function openFeedbackModal() {
    setShowFeedbackModal(true);
  }

  function closeFeedbackModal() {
    setShowFeedbackModal(false);
  }

  function toggleFeedbackModal() {
    setShowFeedbackModal(!showFeedbackModel);
  }

  const screenshotModalFilter = (node: HTMLElement) =>
    !node.id?.startsWith(FEEDBACK_MODAL_ID);

  const FeedbackModal = () => {
    const [email, updateEmail] = useState('');
    const [description, updateDescription] = useState('');
    const [isTakingScreenCapture, setIsTakingScreenCapture] = useState(false);
    const [fileAttachement, setFileAttachement] = useState<File | null>(null);
    const [isSendingFeedback, setIsSendingFeedback] = useState(false);
    const { toast, setToast, dismissToast } = useToast();

    async function takeScreenshot() {
      if (
        document.documentElement === null ||
        isSendingFeedback ||
        isTakingScreenCapture ||
        fileAttachement !== null
      ) {
        return;
      }
      setIsTakingScreenCapture(true);

      toBlob(document.documentElement, {
        cacheBust: true,
        filter: screenshotModalFilter
      })
        .then(async blob => {
          // setScreenshotBlob(blob);
          if (blob !== null) {
            const screenShotFile = new File(
              [blob],
              'Automated screen capture',
              {
                type: blob.type
              }
            );

            setFileAttachement(screenShotFile);
          }

          setToast({
            hasError: false,
            title: 'Screen capture',
            message:
              'A screenshot of your current page on Komiser has been captured and attached to your feedback.'
          });
        })
        .catch(err => {
          setToast({
            hasError: true,
            title: 'Screen capture failed',
            message:
              'The capture of your current page on Komiser couldn’t be saved. Please try again or upload a screenshot manually. Our support is also happy to help you!'
          });
        })
        .finally(() => {
          setIsTakingScreenCapture(false);
        });
    }

    function clearFeedbackForm() {
      setFileAttachement(null);
      updateDescription('');
      updateEmail('');
    }

    async function uploadFeedback(e: SyntheticEvent) {
      if (!isSendingFeedback) {
        try {
          setIsSendingFeedback(true);
          e.preventDefault();
          const formData = new FormData();

          formData.append('description', description);
          if (email) formData.append('email', email);
          if (fileAttachement && fileAttachement !== null)
            formData.append('image', fileAttachement);

          settingsService
            .sendFeedback(formData)
            .then(result => {
              setToast({
                hasError: false,
                title: 'Feedback sent',
                message:
                  result.Response ||
                  'Your insights are valuable to us. Thank you for helping us improve!'
              });
              clearFeedbackForm();
            })
            .catch(error => {
              setToast({
                hasError: true,
                title: 'Feedback',
                message: 'An Error happened. Maybe try again please!'
              });
            })
            .finally(() => {
              setIsSendingFeedback(false);
            });
        } catch {
          setIsSendingFeedback(false);
        }
      }
    }

    function uploadFile(attachement: File) {
      setFileAttachement(attachement);
    }

    return (
      <>
        <Modal
          isOpen={showFeedbackModel}
          closeModal={() => closeFeedbackModal()}
          id={FEEDBACK_MODAL_ID}
        >
          <div className="w-[546px]">
            <h3 className="text-lg font-bold text-black-900">
              Describe your issue
            </h3>
            <p className="text-base text-black-400">
              By providing details of the issue you’ve encountered and outlining
              the steps to reproduce it, we’ll be able to give you better
              support.
            </p>
            <form onSubmit={uploadFeedback} className="mt-4">
              <Input
                disabled={isSendingFeedback}
                type="email"
                label="Your email"
                name="email"
                action={change => {
                  updateEmail(change.email);
                }}
                value={email}
                required
              />
              <textarea
                disabled={isSendingFeedback}
                rows={13}
                placeholder={textAreaPlaceholder}
                className="peer mt-4 w-full rounded bg-white px-4 pb-[0.75rem] pt-[1.75rem] text-sm text-black-900 caret-primary outline outline-[0.063rem] outline-black-200 focus:outline-[0.12rem] focus:outline-primary"
                onChange={event => updateDescription(event?.target?.value)}
                value={description}
                required
              />
              <div className="mt-4 h-full w-full">
                <div className="flex basis-1/2 items-stretch justify-stretch gap-2">
                  {fileAttachement === null && (
                    <div
                      onClick={() => takeScreenshot()}
                      className="w-[50%] grow cursor-pointer rounded border-2 border-black-170 py-5 text-center text-xs transition hover:border-[#B6EAEA] hover:bg-black-100"
                    >
                      <svg
                        className="mb-2 inline-block"
                        width="25"
                        height="24"
                        viewBox="0 0 25 24"
                        fill="none"
                        xmlns="http://www.w3.org/2000/svg"
                      >
                        <path
                          d="M2.25 9V6.5C2.25 4.01 4.26 2 6.75 2H9.25"
                          stroke="#0C1717"
                          stroke-width="1.5"
                          stroke-linecap="round"
                          stroke-linejoin="round"
                        />
                        <path
                          d="M15.25 2H17.75C20.24 2 22.25 4.01 22.25 6.5V9"
                          stroke="#0C1717"
                          stroke-width="1.5"
                          stroke-linecap="round"
                          stroke-linejoin="round"
                        />
                        <path
                          d="M22.25 16V17.5C22.25 19.99 20.24 22 17.75 22H16.25"
                          stroke="#0C1717"
                          stroke-width="1.5"
                          stroke-linecap="round"
                          stroke-linejoin="round"
                        />
                        <path
                          d="M9.25 22H6.75C4.26 22 2.25 19.99 2.25 17.5V15"
                          stroke="#0C1717"
                          stroke-width="1.5"
                          stroke-linecap="round"
                          stroke-linejoin="round"
                        />
                      </svg>

                      <p>
                        {isTakingScreenCapture
                          ? 'Taking screencapture...'
                          : 'Capture current screen'}
                      </p>
                    </div>
                  )}
                  <FileUploader
                    classes={
                      fileAttachement === null
                        ? `grow cursor-pointer rounded border-2 border-dashed border-black-170 py-5 text-center text-xs transition hover:border-[#B6EAEA] hover:bg-black-100 w-[50%]`
                        : ` flex-1 rounded border-2 border-[#B6EAEA] px-6 py-5 text-center text-xs transition w-full`
                    }
                    disabled={
                      fileAttachement !== null ||
                      isSendingFeedback ||
                      isTakingScreenCapture
                    }
                    handleChange={uploadFile}
                    name="attachement"
                    types={FILE_TYPES}
                    fileOrFiles={fileAttachement}
                    hoverTitle="drop here"
                    maxSize={MAX_FILE_SIZE_MB}
                    onTypeError={(err: string) =>
                      setToast({
                        hasError: true,
                        title: 'File upload failed',
                        message: err
                      })
                    }
                    onSizeError={(err: string) =>
                      setToast({
                        hasError: true,
                        title: 'File upload failed',
                        message: err
                      })
                    }
                    dropMessageStyle={{
                      width: '100%',
                      height: '100%',
                      position: 'absolute',
                      background: '#F4F9F9',
                      top: 0,
                      right: 2,
                      display: 'flex',
                      flexGrow: 2,
                      opacity: 1,
                      zIndex: 20,
                      color: '#008484',
                      fontSize: 14,
                      border: 'none'
                    }}
                  >
                    {fileAttachement === null && (
                      <div className="">
                        <svg
                          className="mb-2 inline-block"
                          width="25"
                          height="24"
                          viewBox="0 0 25 24"
                          fill="none"
                          xmlns="http://www.w3.org/2000/svg"
                        >
                          <path
                            d="M9.5 17V11L7.5 13"
                            stroke="#0C1717"
                            stroke-width="1.5"
                            stroke-linecap="round"
                            stroke-linejoin="round"
                          />
                          <path
                            d="M9.5 11L11.5 13"
                            stroke="#0C1717"
                            stroke-width="1.5"
                            stroke-linecap="round"
                            stroke-linejoin="round"
                          />
                          <path
                            d="M22.5 10V15C22.5 20 20.5 22 15.5 22H9.5C4.5 22 2.5 20 2.5 15V9C2.5 4 4.5 2 9.5 2H14.5"
                            stroke="#0C1717"
                            stroke-width="1.5"
                            stroke-linecap="round"
                            stroke-linejoin="round"
                          />
                          <path
                            d="M22.5 10H18.5C15.5 10 14.5 9 14.5 6V2L22.5 10Z"
                            stroke="#0C1717"
                            stroke-width="1.5"
                            stroke-linecap="round"
                            stroke-linejoin="round"
                          />
                        </svg>
                        <p>
                          Drag and drop or{' '}
                          <span className="cursor-pointer text-primary">
                            choose a file
                          </span>
                        </p>
                      </div>
                    )}
                    {fileAttachement !== null && (
                      <div className="relative h-full w-full">
                        <div className="flex h-full w-full items-center justify-items-center gap-2">
                          <svg
                            width="40"
                            height="40"
                            viewBox="0 0 40 40"
                            fill="none"
                            xmlns="http://www.w3.org/2000/svg"
                          >
                            <path
                              d="M35 11.6666V28.3333C35 33.3333 32.5 36.6666 26.6667 36.6666H13.3333C7.5 36.6666 5 33.3333 5 28.3333V11.6666C5 6.66659 7.5 3.33325 13.3333 3.33325H26.6667C32.5 3.33325 35 6.66659 35 11.6666Z"
                              stroke="#0C1717"
                              strokeWidth="1.5"
                              strokeMiterlimit="10"
                              strokeLinecap="round"
                              strokeLinejoin="round"
                            />
                            <path
                              d="M24.167 7.5V10.8333C24.167 12.6667 25.667 14.1667 27.5003 14.1667H30.8337"
                              stroke="#0C1717"
                              strokeWidth="1.5"
                              strokeMiterlimit="10"
                              strokeLinecap="round"
                              strokeLinejoin="round"
                            />
                            <path
                              d="M13.333 21.6667H19.9997"
                              stroke="#0C1717"
                              strokeWidth="1.5"
                              strokeMiterlimit="10"
                              strokeLinecap="round"
                              strokeLinejoin="round"
                            />
                            <path
                              d="M13.333 28.3333H26.6663"
                              stroke="#0C1717"
                              strokeWidth="1.5"
                              strokeMiterlimit="10"
                              strokeLinecap="round"
                              strokeLinejoin="round"
                            />
                          </svg>
                          <div className="flex-1 text-left">
                            <p>{fileAttachement.name}</p>
                            <p className="text-black-400">
                              {(fileAttachement.size / (1024 * 1024)).toFixed(
                                2
                              )}
                              MB
                            </p>
                          </div>
                        </div>
                        <a
                          onClick={e => {
                            e.preventDefault();
                            if (!isSendingFeedback && !isTakingScreenCapture)
                              setFileAttachement(null);
                            return false;
                          }}
                          className="absolute right-4 top-4 block h-4 w-4 cursor-pointer"
                          aria-disabled={
                            isTakingScreenCapture || isSendingFeedback
                          }
                        >
                          <svg
                            xmlns="http://www.w3.org/2000/svg"
                            width="16"
                            height="17"
                            viewBox="0 0 16 17"
                            fill="none"
                          >
                            <path
                              d="M4.66602 12.079L11.3327 5.41235"
                              stroke="#0C1717"
                              strokeLinecap="round"
                              strokeLinejoin="round"
                            />
                            <path
                              d="M11.3327 12.079L4.66602 5.41235"
                              stroke="#0C1717"
                              strokeLinecap="round"
                              strokeLinejoin="round"
                            />
                          </svg>
                        </a>
                      </div>
                    )}
                  </FileUploader>
                </div>
              </div>
              <div className="mt-4 flex justify-between">
                <p className="text-xs text-black-400">
                  Need in depth assistance?
                  <br />
                  Email us at{' '}
                  <a
                    href="mailto:support@tailwarden.com"
                    className="text-primary"
                  >
                    support@tailwarden.com
                  </a>
                  .
                </p>
                <Button type="submit" size="xs" disabled={isSendingFeedback}>
                  {isSendingFeedback ? `Sending...` : `Send Feedback`}
                </Button>
              </div>
            </form>
          </div>
        </Modal>
        {toast && <Toast {...toast} dismissToast={dismissToast} />}
      </>
    );
  };

  return {
    openFeedbackModal,
    closeFeedbackModal,
    toggleFeedbackModal,
    FeedbackModal: memo(FeedbackModal)
  };
};

export default useFeedbackWidget;
