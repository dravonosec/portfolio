package Java.Logic.Socket;

import io.netty.channel.Channel;
import io.netty.channel.ChannelHandlerContext;

import java.util.function.Consumer;

public class AxServerHandler extends AxAbstractHandler {
    private static Channel serverChannel;

    public AxServerHandler(Consumer<String> messageCallback) {
        super(messageCallback);
    }


    @Override
    public void channelActive(ChannelHandlerContext ctx) {
        // При подключении клиента сохарнить соединение
        ctx.channel().writeAndFlush(Utils.stringToByteBuf("connectionIsActive"));
        serverChannel = ctx.channel();
    }

    @Override
    public void channelInactive(ChannelHandlerContext ctx) {
        // При отключении обнулить соединение
        serverChannel = null;
    }

    @Override
    public void channelRead(ChannelHandlerContext ctx, Object msg) {
        // Логика обработки входящих сообщений от клиента
        String message = Utils.objectToString(msg);
        messageCallback.accept(message);
    }

    // Отправка сообщений клиенту
    @Override
    public void channelWrite(String message) {
        serverChannel.writeAndFlush(Utils.stringToByteBuf(message));
    }

    @Override
    public boolean checkConnection(){
        return AxServer.isActive();
    }

}

