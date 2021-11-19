import { NextPage } from 'next'
import { Message, useMessagePostedSubscription, useMessagesQuery, usePostMessageMutation } from "@/generated/graphql";
import { useRouter } from "next/router";
import { useCallback, useEffect, useState } from "react";

const Index: NextPage = () => {

    // const { roomId } = useRouter()
    const roomId = "0512cf5a-f822-4904-8117-aab8e2f9e4aa"

    const [messages, setMessages] = useState<Message[]>([]);
    const [inputValue, setInputValue] = useState('')
    const { loading, data, error } = useMessagesQuery({
        variables: {
            roomId: roomId
        }
    })
    const [createMessage] = usePostMessageMutation()

    useEffect(() => {
        if (data?.messages) setMessages(data.messages)
    }, [data?.messages])

    const { data: subData } = useMessagePostedSubscription({
      variables: {
        roomId: roomId,
      },
        onSubscriptionData: (data) => {
            console.log(data)
        }
    })

  useEffect(() => {
    if (subData?.messagePosted) setMessages(m => [...m, subData?.messagePosted])
  }, [subData?.messagePosted])

    const handleChange = useCallback((e: React.ChangeEvent<HTMLInputElement>) => {
        setInputValue(e.target.value)
    }, [])

  const handleClick = useCallback(async (e) => {
    e.preventDefault()
    await createMessage({ variables: { roomId: roomId, text: inputValue } })
  }, [inputValue, createMessage])

    if (data?.messages.length == 0) {
        return <div>No data found</div>
    }

    return (
      <div className="flex flex-col h-screen bg-gray-50">
        <div className="grid place-items-center mx-2 md:my-20 my-10">
          <div
            className="w-11/12 p-12 sm:w-8/12 md:w-6/12 lg:w-5/12 2xl:w-4/12
            px-6 py-10 sm:px-10 sm:py-6
            bg-white rounded-lg shadow-md lg:shadow-lg"
          >
            <h2 className="text-center font-semibold text-3xl lg:text-4xl text-gray-800">
              Chat
            </h2>
            {messages.map((m, i) => <div key={i}>{m.text}</div>)}
            <input
              value={inputValue}
              onChange={handleChange}
              className="block w-full py-3 px-1 mt-2
                      text-gray-800 appearance-none
                      border-b-2 border-gray-100
                      focus:text-gray-500 focus:outline-none focus:border-gray-200"
            />
            <button onClick={handleClick}>Submit</button>
          </div>
        </div>
      </div>
    )
}

export default Index;
