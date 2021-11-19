import { Message } from "@/generated/graphql";

type MessagesProps = {
  messages: Message[]
}

const Messages = ({ messages }: MessagesProps) => (
  <div id="messages" className="flex flex-col space-y-4 p-3 overflow-y-auto">
    { messages.map((m, i) => (
      <MessageDetail key={i} message={m} />
    ))}
    <div className="chat-message">
      <div className="flex items-end">
        <div className="flex flex-col space-y2 text-xs max-w-xs mx-2 order-2 items-start items-start">
                <span className="px-4 py-2 rounded-lg inline-block rounded-bl-none
                  bg-gray-300 text-gray-600">
                  Can be verified on aby platform using docker?
                </span>
        </div>
      </div>
    </div>
    <div className="chat-message">
      <div className="flex items-end justify-end">
        <div className="flex flex-col space-y-2 text-xs max-w-xs mx-2 order-1 items-end">
          <span className="px-4 py-2 rounded-lg inline-block rounded-br-none
            bg-blue-600 text-white">
            Your error message says permission denied, npm global installs must be given root privileges.
          </span>
        </div>
      </div>
    </div>
  </div>
)

type MessageProps = {
  message: Message
}

const MessageDetail = ({ message }:MessageProps) => {
  return (
    <div className="chat-message">
      {/* justify-end を付加する */}
      <div className="flex items-end">
        {/*　items-start => items-end　*/}
        <div className="flex flex-col space-y2 text-xs max-w-xs mx-2 order-2 items-start">
          {/*テキストと背景をかえる bg-gray-300 text-gray-600 -> bg-blue-600 text-white */}
          <span className="px-4 py-2 rounded-lg inline-block rounded-bl-none
            bg-gray-300 text-gray-600">
            {message.text}
          </span>
        </div>
      </div>
    </div>
  )
}

export default Messages
