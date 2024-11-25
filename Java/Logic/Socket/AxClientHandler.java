package Java.Logic.Socket;

import io.netty.channel.Channel;
import io.netty.channel.ChannelHandlerContext;

import java.util.function.Consumer;


public class AxClientHandler extends AxAbstractHandler {
    private static Channel clientChannel;


    public AxClientHandler(Consumer<String> messageCallback) {
        super(messageCallback);
    }

    @Override
    public void channelActive(ChannelHandlerContext ctx) {
        // При подключении клиента сохранить соединение
        ctx.channel().writeAndFlush(Utils.stringToByteBuf("connectionIsActive"));
        clientChannel = ctx.channel();
    }

    @Override
    public void channelInactive(ChannelHandlerContext ctx) {
        clientChannel = null;
    }

    // Чтение канала и получение сообщений
    @Override
    public void channelRead(ChannelHandlerContext ctx, Object msg) {
        String message = Utils.objectToString(msg);
        messageCallback.accept(message);
    }

    // Отправка сообщений серверу
    @Override
    public void channelWrite(String message) {
        clientChannel.writeAndFlush(Utils.stringToByteBuf(message));
    }

    @Override
    public boolean checkConnection() {
        return AxClient.isActive();
    }


    @Override
    public boolean isActive() {
        return clientChannel.isActive();
    }
}

