package coders;

import java.awt.image.BufferedImage;
import java.io.ByteArrayOutputStream;
import java.io.IOException;
import java.util.Base64;

import javax.imageio.ImageIO;
import javax.json.Json;
import javax.websocket.EncodeException;
import javax.websocket.Encoder;
import javax.websocket.EndpointConfig;

public class MessageEncoder implements Encoder.Text< Message > {
	@Override
	public void init( final EndpointConfig config ) {
	}

	@Override
	public String encode( final Message message ) throws EncodeException {
		String encodeMessage = null;
		if(message.getAction()=="IMG_FRAME"){
			encodeMessage=encodeImage(message);
		}else if(message.getAction()=="MOUSE_MOVE"){
			encodeMessage=encodeMoveMouse(message);
		}else if(message.getAction()=="MOUSE_RCLICK" || message.getAction()=="MOUSE_LCLICK"){
			encodeMessage=encodeMouseClick(message);
		}
		return encodeMessage;
	}
	
	private String encodeMouseClick(Message message){
		return Json.createObjectBuilder()
				.add( "action", message.getAction() )
				.build().toString();
	}
	
	private String encodeMoveMouse(Message message){
		return Json.createObjectBuilder()
				.add("data", (String)message.getData())
				.add( "action", message.getAction() )
				.build().toString();
	}

	private String encodeImage(Message message){
		String imageString = null;
		ByteArrayOutputStream baos = new ByteArrayOutputStream();
		try {
			ImageIO.write((BufferedImage)message.getData(),"jpg", baos);
			byte[] imageAsRawBytes = baos.toByteArray();
			imageString = new String(Base64.getEncoder().encode(imageAsRawBytes));
		} catch (IOException ex) {
		}

		return Json.createObjectBuilder()
				.add("data", imageString)
				.add( "action", message.getAction() )
				.build()
				.toString();
	}

	@Override
	public void destroy() {
	}
}
