package Java.Logic.Socket;

import io.netty.channel.Channel;
import io.netty.channel.ChannelHandlerContext;
import io.netty.channel.ChannelInboundHandlerAdapter;

import java.util.function.Consumer;

public abstract class AxAbstractHandler extends ChannelInboundHandlerAdapter {
    Channel abstarctChannel;

    protected final Consumer<String> messageCallback;


    public AxAbstractHandler(Consumer<String> messageCallback) {
        this.messageCallback = messageCallback;
    }

    @Override
    public void channelActive(ChannelHandlerContext ctx) {
        // При подключении сохранить канал, по которому установлено соединение
        ctx.channel().writeAndFlush(Utils.stringToByteBuf("Проверка"));
        this.abstarctChannel = ctx.channel();
    }
    @Override
    public void channelInactive(ChannelHandlerContext ctx) {
        abstarctChannel = null;
    }

    @Override
    public void channelRead(ChannelHandlerContext ctx, Object msg) {
        String message = Utils.objectToString(msg);
        messageCallback.accept(message);
    }

    public void channelWrite(String message) {
        abstarctChannel.writeAndFlush(Utils.stringToByteBuf(message));
    }

    public boolean checkConnection() {
        return false;
    }

    public boolean isActive() {
        return abstarctChannel.isActive();
    }
}

