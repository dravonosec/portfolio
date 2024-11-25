package Java.Logic.Socket;

import io.netty.bootstrap.Bootstrap;
import io.netty.channel.*;
import io.netty.channel.nio.NioEventLoopGroup;
import io.netty.channel.socket.SocketChannel;
import io.netty.channel.socket.nio.NioSocketChannel;

public final class AxClient {
    private final AxClientHandler clientHandler;
    static boolean isActive = false;
    EventLoopGroup group;
    ChannelFuture channelFuture;
    public AxClient(AxClientHandler clientHandler) {
        this.clientHandler = clientHandler;
    }

    public void connectToServer(int port, String hostIp) throws Exception {
        group = new NioEventLoopGroup();

        try {
            Bootstrap b = new Bootstrap();
            b.group(group)
                    .channel(NioSocketChannel.class)
                    .option(ChannelOption.TCP_NODELAY, true)
                    .handler(new ChannelInitializer<SocketChannel>() {
                        @Override
                        public void initChannel(SocketChannel ch) throws Exception {
                            ChannelPipeline p = ch.pipeline();
                            p.addLast(clientHandler);
                        }
                    });
            // Подключение клиента
            isActive = true;
            channelFuture = b.connect(hostIp, port).sync();

            // Ожидание, пока соеинение не будет закрыто
        } finally {
            System.out.println("Connection set");
        }
    }

    public void shutdown() {
        System.out.println("Stopping client");
        isActive = false;
        try {
            group.shutdownGracefully().sync();
            channelFuture.channel().closeFuture().sync();
            System.out.println("Client stopped");
        }
        catch (InterruptedException e) {
            e.printStackTrace();
        }
    }

    public static boolean isActive() {
        return isActive;
    }
}

