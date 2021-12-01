import { Message } from "@/generated/graphql";
import classNames from "classnames";

type MessagesProps = {
  messages: Message[]
}

const Messages = ({ messages }: MessagesProps) => (
  <div id="messages" className="flex flex-col space-y-4 p-3 overflow-y-auto">
    { messages.map((m, i) => (
      <MessageDetail key={i} message={m} />
    ))}
  </div>
)

type MessageProps = {
  message: Message
}

const MessageDetail = ({ message }:MessageProps) => {
  return (
    <div className="chat-message">
      <div className={classNames(
        [
          "flex items-end",
          message.isSender && "justify-end"
        ]
      )}>
        {/*　items-start => items-end　*/}
        <div className={classNames(
          [
            "flex flex-col space-y2 text-xs max-w-xs mx-2 order-2",
            message.isSender ? "justify-end" : "items-start"
          ])}>
          {/*テキストと背景をかえる bg-gray-300 text-gray-600 -> bg-blue-600 text-white */}
          <span className={classNames([
            "px-4 py-2 rounded-lg inline-block rounded-bl-none",
            message.isSender ? "bg-blue-600 text-white" : "bg-gray-300 text-gray-600"])}>
            {message.text}
          </span>
        </div>
      </div>
    </div>
  )
}

export default Messages
