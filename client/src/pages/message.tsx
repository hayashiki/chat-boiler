import { NextPage } from "next";
import Header from "@/components/header";
import Messages from "@/components/messages";
import MessageForm from "@/components/form";
import { Message } from "@/generated/graphql";

const messages = [
  {
    id: "1",
    roomId: "1",
    text: "ラーメン食べたい",
    userID: "taro",
    isSender: true
  },
  {
    id: "2",
    roomId: "2",
    text: "OK,どこの店にしよう？",
    userID: "jiro",
    isSender: false
  }
] as Message[]


const MessagePage: NextPage = () => {
  return (
    <div className="flex-1 sm:p-6 justify-between flex flex-col h-screen">
      {/* header area */}
      <Header />
      <Messages messages={messages} />
      <MessageForm />
    </div>
  )
}

export default MessagePage
