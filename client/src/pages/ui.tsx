import {FC} from "react";
import GoogleSvg from "@/images/Google.svg"
import Bookmark from "@/images/bookmarks_black_24dp.svg"
import Dashboard from "@/images/dashboard_black_24dp.svg"
import CheckCircle from "@/images/check_circle_black_24dp.svg"
import SearchSVG from "@/images/search_black_24dp.svg"
import Image from 'next/image'

const ChatUI = () => {
  return (
    <div className="flex">
      {/*w-12	width: 3rem;*/}
      {/*Use h-screen to make an element span the entire height of the viewport.*/}
      {/* justify-betweenでアイコングループたちを等間隔にする */}
      <div className="w-14 h-screen bg-green-100 flex flex-col items-center justify-between">
        <div className="pt-3">
          <GoogleSvg />
        </div>

        <div className="flex flex-col">
          <div className="py-3 w-5">
            <svg xmlns="http://www.w3.org/2000/svg" height="24px" viewBox="0 0 24 24" width="24px" fill="#000000"><path d="M0 0h24v24H0V0z" fill="none"/><path d="M19 18l2 1V3c0-1.1-.9-2-2-2H8.99C7.89 1 7 1.9 7 3h10c1.1 0 2 .9 2 2v13zM15 5H5c-1.1 0-2 .9-2 2v16l7-3 7 3V7c0-1.1-.9-2-2-2z"/></svg>
          </div>
          <div className="py-3 w-5">
            <Dashboard />
          </div>
          <div className="py-3 w-5">
            <Bookmark  />
          </div>
          <div className="py-3 w-5 flex flex-col justify-center items-center">
            <Dashboard  />
            <div className="w-1 h-1 bg-blue-500 rounded-full"></div>
          </div>
          <div className="py-3 w-5">
            <Bookmark  />
          </div>
        </div>
        <CheckCircle />
      </div>
      {/* ルーム一覧エリア */}
      <div className="w-68 h-screen bg-gray-100">
        <div className="text-xl p-3">
          Chats
        </div>
        <div className="p-3 flex">
          <input
            // inputのp-2は入力ゾーンがおおきくなる
            className="p-2 w-10/12 bg-gray-200 text-xs focus:outline-none rounded-tl-md rounded-bl-md"
            type="text" placeholder="Search for messages or users ..."/>
          <div className="p-2 w-2/12 flex justify-center items-center bg-gray-200 rounded-tr-md rounded-br-md">
            <SearchSVG className="w-4"/>
          </div>
        </div>
        <div className="flex">
          <div className="p-2 flex flex-col">
            <div className="w-8 h-8 relative">
              <Image
                className="rounded-full"
                layout="fill"
                src="/avatar.png"
                alt="me" />
            </div>
            <div className="text-gray-500 text-xs pt-1 text-center">
              William
            </div>
          </div>

          <div className="p-2 flex flex-col">
            <div className="w-8 h-8 relative">
              <Image
                className="rounded-full"
                layout="fill"
                src="/avatar.png"
                alt="me" />
            </div>
            <div className="text-gray-500 text-xs pt-1 text-center">
              Jone
            </div>
          </div>

        </div>

        <div className="flex m-3 bg-white rounded-lg p-2">
          <div>
            <div className="w-14 h-14 relative">
              <Image
                className="rounded-full"
                layout="fill"
                src="/avatar.png"
                alt="me" />
            </div>
          </div>
          <div className="flex-grow p-3">
            <div className="flex text-xs justify-between">
              <div>William</div>
              <div className="text-gray-400">12:00 AM</div>
            </div>
            <div className="text-xs text-gray-500">this is really dope and i am very ..</div>
          </div>
        </div>
      </div>
      {/* 残り flex-grow */}
      <div className="flex-grow h-screen flex flex-col">
        <div className="w-full h-14 flex justify-between">
          <div className="flex items-center">
            <div className="p-3">
              <div className="w-8 h-8 relative">
                <Image
                  layout="fill"
                  src="/avatar.png"
                  className="rounded-full avatar avatar-user width-full border color-bg-default"
                />
              </div>
              <div className="justify-center items-center w-3 h-3 relative left-6">
                <div className="w-2 h-2 bg-green-600 rounded-full"></div>
              </div>
            </div> {/* p-3 */}

            <div className="p-3">
              <div className="text-sm">William</div>
              <div className="text-xs">online</div>
            </div>
          </div>

          {/* 右エリア サーチと、モア*/}
          <div className="flex items-center p-3">
            <div className="w-5 g-5">
              <Bookmark  />
            </div>
            <div className="w-5 g-5">
              <Bookmark  />
            </div>
          </div>

        </div>
        {/* 中段エリア */}
        <div className="w-full flex-grow bg-blue-100">
          <div className="flex items-end w-3/6 bg-gray-100 m-8">
            <div className="w-8 h-8 relative m-3">
              <Image
                layout="fill"
                src="/avatar.png"
                className="rounded-full avatar avatar-user width-full border color-bg-default"
              />
            </div>
            <div className="p-3">
              <div className="text-sm">William James</div>
              <div className="text-xs text-gray-400">
                Lorem ipsum dolor sit amet consecurter
              </div>
              <div className="text-xs text-gray-400">
                8 minutes ago
              </div>
            </div>

          </div>

        </div>
        <div className="w-full h-14 bg-red-100">abc</div>
      </div>
    </div>
  )
}

export default ChatUI
