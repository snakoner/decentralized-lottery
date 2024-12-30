// ChatRoom.tsx
import React, { useState } from 'react';
import { List, Input, Button, Typography } from 'antd';
import { SendOutlined } from '@ant-design/icons';

const { Text } = Typography;

interface Message {
    id: number;
    date: number;
    user: string;
    content: string;
}

const ChatRoom: React.FC = (account?: string) => {
    const [messages, setMessages] = useState<Message[]>([]);
    const [inputValue, setInputValue] = useState<string>('');
    const ws = new WebSocket('ws://localhost:8000/ws');

    const handleSend = () => {
        if (inputValue.trim()) {
            const newMessage: Message = {
                id: 0,
                date: Date.now(),
                user: 'You',
                content: inputValue.trim(),
            };

            if (ws) {
                try {
                    ws.send(JSON.stringify(newMessage));
                } catch(error) {
                    console.log(error);
                }
            }
            setMessages([...messages, newMessage]);
            setInputValue('');
        }
    };

    ws.onopen = () => {
        console.log('Connected to server');
        // ws.send('Hello, server!');
    };

    ws.onmessage = (event) => {
        console.log('Received:', event.data);
    };

    ws.onclose = () => {
        console.log('Disconnected from server');
    };

    ws.onerror = (error) => {
        console.error('WebSocket error:', error);
    };


  return (
    <div style={{ maxWidth: 960, margin: '0 auto', padding: '20px' }}>
      <List
        header={<div>Participants chat</div>}
        bordered
        dataSource={messages}
        renderItem={(item) => (
          <List.Item key={item.id}>
            <Text strong>{item.user}:</Text> {item.content}
          </List.Item>
        )}
        style={{ marginBottom: '20px', backgroundColor: '#fff' }}
      />
      <Input.Group compact>
        <Input
          style={{ width: 'calc(100% - 50px)' }}
          value={inputValue}
          onChange={(e) => setInputValue(e.target.value)}
          onPressEnter={handleSend}
          placeholder="Enter your message..."
        />
        <Button
          type="primary"
          icon={<SendOutlined />}
          onClick={handleSend}
        />
      </Input.Group>
    </div>
  );
};

export default ChatRoom;